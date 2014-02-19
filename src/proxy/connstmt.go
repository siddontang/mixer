package proxy

import (
	"encoding/binary"
	"fmt"
	"lib/log"
	"math"
	. "mysql"
	. "parser"
	"strconv"
)

type stmt struct {
	id uint32
	l  *lex

	params  []Field
	columns []Field
}

func (s *stmt) ResetParams() {
	s.l.Args = make([]interface{}, len(s.params))
}

func routePrepare(nodeName string, co *Conn, query string, args ...interface{}) interface{} {
	if len(args) > 0 {
		return fmt.Errorf("prepare cannot have args")
	}

	r, err := co.Prepare(query)
	if err != nil {
		log.Error("node %s exec error %s", nodeName, err.Error())
		return err
	} else {
		r.Close()
		return [][]Field{r.Params, r.Columns}
	}
}

func (c *conn) handleStmtPrepare(data []byte) error {
	s := new(stmt)
	query := string(data)

	l, err := parseQuery(query)
	if err != nil {
		return err
	}

	l.Prepared = true
	s.l = l

	var results []interface{}
	results, err = c.route(l, routePrepare)
	if err != nil {
		return err
	}

LOOP:
	for _, i := range results {
		switch v := i.(type) {
		case error:
			err = v
			break LOOP
		case ([][]Field):
			if len(v) != 2 {
				err = fmt.Errorf("invalid prepare result %d", len(v))
			}
			//now we only use columns and params
			s.params = v[0]
			s.columns = v[1]
		default:
			err = fmt.Errorf("invalid return type %T", i)
			break LOOP
		}
	}

	if err != nil {
		return err
	}

	s.id = c.stmtId
	c.stmtId++

	if err = c.writePrepare(s); err != nil {
		return err
	}

	s.ResetParams()

	c.stmts[s.id] = s

	return nil
}

func (c *conn) writePrepare(s *stmt) error {
	data := make([]byte, 4, 128)

	//status ok
	data = append(data, 0)
	//stmt id
	data = append(data, Uint32ToBytes(s.id)...)
	//number columns
	data = append(data, Uint16ToBytes(uint16(len(s.columns)))...)
	//number params
	data = append(data, Uint16ToBytes(uint16(len(s.params)))...)
	//filter [00]
	data = append(data, 0)
	//warning count
	data = append(data, 0, 0)

	if err := c.WritePacket(data); err != nil {
		return err
	}

	if len(s.params) > 0 {
		for _, v := range s.params {
			data = data[0:4]
			data = append(data, v.Dump()...)
			if err := c.WritePacket(data); err != nil {
				return err
			}
		}

		if err := c.writeEOF(c.status); err != nil {
			return err
		}
	}

	if len(s.columns) > 0 {
		for _, v := range s.columns {
			data = data[0:4]
			data = append(data, v.Dump()...)
			if err := c.WritePacket(data); err != nil {
				return err
			}
		}

		if err := c.writeEOF(c.status); err != nil {
			return err
		}

	}
	return nil
}

func (c *conn) handleStmtExecute(data []byte) error {
	if len(data) < 9 {
		return ErrMalformPacket
	}

	pos := 0
	id := binary.LittleEndian.Uint32(data[0:4])
	pos += 4

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_execute")
	}

	flag := data[pos]
	pos++
	//now we only support CURSOR_TYPE_NO_CURSOR flag
	if flag != 0 {
		return NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("unsupported flag %d", flag))
	}

	//skip iteration-count, always 1
	pos += 4

	var nullBitmaps []byte
	var paramTypes []byte
	var paramValues []byte

	paramNum := len(s.params)

	if paramNum > 0 {
		nullBitmapLen := (len(s.params) + 7) >> 3
		if len(data) < (pos + nullBitmapLen + 1) {
			return ErrMalformPacket
		}
		nullBitmaps = data[pos : pos+nullBitmapLen]
		pos += nullBitmapLen

		//new param bound flag
		if data[pos] == 1 {
			pos++
			if len(data) < (pos + (paramNum << 1)) {
				return ErrMalformPacket
			}

			paramTypes = data[pos : pos+(paramNum<<1)]
			pos += (paramNum << 1)

			paramValues = data[pos:]
		}

		if err := c.bindStmtArgs(s, nullBitmaps, paramTypes, paramValues); err != nil {
			return err
		}
	}

	var err error
	switch s.l.Get(0).Type {
	case TK_SQL_SELECT:
		err = c.handleSelect(s.l)
	case TK_SQL_INSERT:
		err = c.handleExec(s.l)
	case TK_SQL_UPDATE:
		err = c.handleExec(s.l)
	case TK_SQL_DELETE:
		err = c.handleExec(s.l)
	case TK_SQL_REPLACE:
		err = c.handleExec(s.l)
	default:
		err = NewError(ER_UNKNOWN_ERROR,
			fmt.Sprintf("command %s not supported now", s.l.Get(0).Value))
	}

	s.l.Args = make([]interface{}, paramNum)

	return err
}

func (c *conn) bindStmtArgs(s *stmt, nullBitmap, paramTypes, paramValues []byte) error {
	args := s.l.Args

	pos := 0

	var v []byte
	var n int = 0
	var isNull bool
	var err error

	for i := 0; i < len(s.params); i++ {
		if nullBitmap[i>>3]&(1<<(uint(i)%8)) > 0 {
			args[i] = nil
			continue
		}

		tp := paramTypes[i<<1]
		isUnsigned := (paramTypes[(i<<1)+1] & 0x80) > 0

		switch tp {
		case MYSQL_TYPE_NULL:
			args[i] = nil
			continue

		case MYSQL_TYPE_TINY:
			if len(paramValues) < (pos + 1) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint8(paramValues[pos])
			} else {
				args[i] = int8(paramValues[pos])
			}

			pos++
			continue

		case MYSQL_TYPE_SHORT, MYSQL_TYPE_YEAR:
			if len(paramValues) < (pos + 2) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint16(binary.LittleEndian.Uint16(paramValues[pos : pos+2]))
			} else {
				args[i] = int16((binary.LittleEndian.Uint16(paramValues[pos : pos+2])))
			}
			pos += 2
			continue

		case MYSQL_TYPE_INT24, MYSQL_TYPE_LONG:
			if len(paramValues) < (pos + 4) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint32(binary.LittleEndian.Uint32(paramValues[pos : pos+4]))
			} else {
				args[i] = int32(binary.LittleEndian.Uint32(paramValues[pos : pos+4]))
			}
			pos += 4
			continue

		case MYSQL_TYPE_LONGLONG:
			if len(paramValues) < (pos + 8) {
				return ErrMalformPacket
			}

			if isUnsigned {
				args[i] = binary.LittleEndian.Uint64(paramValues[pos : pos+8])
			} else {
				args[i] = int64(binary.LittleEndian.Uint64(paramValues[pos : pos+8]))
			}
			pos += 8
			continue

		case MYSQL_TYPE_FLOAT:
			if len(paramValues) < (pos + 4) {
				return ErrMalformPacket
			}

			args[i] = float32(math.Float32frombits(binary.LittleEndian.Uint32(paramValues[pos : pos+4])))
			pos += 4
			continue

		case MYSQL_TYPE_DOUBLE:
			if len(paramValues) < (pos + 8) {
				return ErrMalformPacket
			}

			args[i] = math.Float64frombits(binary.LittleEndian.Uint64(paramValues[pos : pos+8]))
			pos += 8
			continue

		case MYSQL_TYPE_DECIMAL, MYSQL_TYPE_NEWDECIMAL, MYSQL_TYPE_VARCHAR,
			MYSQL_TYPE_BIT, MYSQL_TYPE_ENUM, MYSQL_TYPE_SET, MYSQL_TYPE_TINY_BLOB,
			MYSQL_TYPE_MEDIUM_BLOB, MYSQL_TYPE_LONG_BLOB, MYSQL_TYPE_BLOB,
			MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING, MYSQL_TYPE_GEOMETRY,
			MYSQL_TYPE_DATE, MYSQL_TYPE_NEWDATE,
			MYSQL_TYPE_TIMESTAMP, MYSQL_TYPE_DATETIME, MYSQL_TYPE_TIME:
			if len(paramValues) < (pos + 1) {
				return ErrMalformPacket
			}

			v, isNull, n, err = LengthEnodedString(paramValues[pos:])
			pos += n
			if err != nil {
				return err
			}

			if !isNull {
				args[i] = v
				continue
			} else {
				args[i] = nil
				continue
			}
		default:
			return fmt.Errorf("Stmt Unknown FieldType %d", tp)
		}
	}
	return nil
}

func (c *conn) handleStmtSendLongData(data []byte) error {
	if len(data) < 6 {
		return ErrMalformPacket
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_send_longdata")
	}

	paramId := binary.LittleEndian.Uint16(data[4:6])
	if paramId >= uint16(len(s.params)) {
		return NewDefaultError(ER_WRONG_ARGUMENTS, "stmt_send_longdata")
	}

	if s.l.Args[paramId] == nil {
		s.l.Args[paramId] = data[6:]
	} else {
		if b, ok := s.l.Args[paramId].([]byte); ok {
			b = append(b, data[6:]...)
			s.l.Args[paramId] = b
		} else {
			return NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("invalid param long data type %T", s.l.Args[paramId]))
		}
	}

	return nil
}

func (c *conn) handleStmtReset(data []byte) error {
	if len(data) < 4 {
		return ErrMalformPacket
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_reset")
	}

	s.ResetParams()

	return c.writeOK(nil)
}

func (c *conn) handleStmtClose(data []byte) error {
	if len(data) < 4 {
		log.Error("stmt close error %s", ErrMalformPacket.Error())
		return nil
	}

	id := binary.LittleEndian.Uint32(data[0:4])

	_, ok := c.stmts[id]
	if !ok {
		err := NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_close")
		log.Error("stmt close error %s", err.Error())
	} else {
		delete(c.stmts, id)
	}

	return nil
}
