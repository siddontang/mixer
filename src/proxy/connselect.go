package proxy

import (
	"bytes"
	"errors"
	"fmt"
	"lib/log"
	. "mysql"
	. "parser"
	"strconv"
	"strings"
)

func (c *conn) writeResultset(r *Resultset) error {
	r.Status |= c.status

	c.affectedRows = int64(-1)

	columnLen := PutLengthEncodedInt(uint64(len(r.Fields)))

	data := make([]byte, 4, 1024)

	data = append(data, columnLen...)
	if err := c.WritePacket(data); err != nil {
		return err
	}

	for _, v := range r.Fields {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := c.WritePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(r.Status); err != nil {
		return err
	}

	for _, v := range r.RowPackets {
		data = data[0:4]
		data = append(data, v...)
		if err := c.WritePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(r.Status); err != nil {
		return err
	}

	return nil
}

func (c *conn) handleSelectLastInsertId(l *lex) (*Resultset, error) {
	//only support select last_insert_id();
	//not support select last_insert_id(expr);

	if !(l.Get(2).Type == TK_LPAREN && l.Get(3).Type == TK_RPAREN) {
		return nil, NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("only select last_insert_id(); supported now"))
	}

	name := l.Get(1).Value
	if l.Get(4).Type == TK_SQL_AS {
		name = l.Get(5).Value
		if len(name) == 0 {
			return nil, NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("some error around as %s", l.Get(5).Value))
		}
	}

	r := new(Resultset)

	field := Field{}

	field.Name = []byte(name)
	field.OrgName = []byte(l.Get(1).Value)

	//see in wireshake
	field.Charset = 63

	field.Type = MYSQL_TYPE_LONGLONG
	//after mysql 5.6.9, last_insert_id is unsigned, but now we use signed
	field.Flag = BINARY_FLAG | NOT_NULL_FLAG

	r.Fields = []Field{field}

	lastId := []byte(strconv.FormatInt(c.lastInsertId, 10))

	r.RowPackets = append(r.RowPackets, PutLengthEncodedString(lastId))

	return r, nil
}

func (c *conn) handleSelectRowCount(l *lex) (*Resultset, error) {
	if !(l.Get(2).Type == TK_LPAREN && l.Get(3).Type == TK_RPAREN) {
		return nil, NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("only select row_count(); supported now"))
	}

	name := l.Get(1).Value
	if l.Get(4).Type == TK_SQL_AS {
		name = l.Get(5).Value
		if len(name) == 0 {
			return nil, NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("some error around as %s", l.Get(5).Value))
		}
	}

	r := new(Resultset)

	field := Field{}

	field.Name = []byte(name)
	field.OrgName = []byte(l.Get(1).Value)

	//see in wireshake
	field.Charset = 63

	field.Type = MYSQL_TYPE_LONGLONG
	//after mysql 5.6.9
	field.Flag = BINARY_FLAG | NOT_NULL_FLAG

	r.Fields = []Field{field}

	rowCount := []byte(strconv.FormatInt(c.affectedRows, 10))

	r.RowPackets = append(r.RowPackets, PutLengthEncodedString(rowCount))

	return r, nil
}

func routeSelect(nodeName string, co *Conn, query string, args ...interface{}) interface{} {
	r, err := co.Query(query, args...)
	if err != nil {
		log.Error("node %s query error %s", nodeName, err.Error())
		return err
	} else {
		return r
	}
}

func mergeResultset(dest *Resultset, src *Resultset) error {
	if dest.ColumnNumber() != src.ColumnNumber() {
		return errors.New("column not match")
	}

	for i := range dest.Fields {
		//here we test name, type and flag
		if !bytes.Equal(dest.Fields[i].Name, src.Fields[i].Name) {
			return fmt.Errorf("field name %s != %s", dest.Fields[i].Name, src.Fields[i].Name)
		}

		if dest.Fields[i].Type != src.Fields[i].Type {
			return fmt.Errorf("field type %d != %d", dest.Fields[i].Type, src.Fields[i].Type)
		}

		if dest.Fields[i].Flag != src.Fields[i].Flag {
			return fmt.Errorf("field flag %d != %d", dest.Fields[i].Flag, src.Fields[i].Flag)
		}
	}

	dest.Status |= src.Status

	//later we may merge with select condition like limit, order by, etc...
	//now we only append row
	for _, v := range src.RowPackets {
		dest.RowPackets = append(dest.RowPackets, v)
	}

	return nil
}

func (c *conn) handleSelectDefault(l *lex) (*Resultset, error) {

	results, err := c.route(l, routeSelect)
	if err != nil {
		return nil, err
	}

	var r *Resultset = nil

LOOP:
	for _, i := range results {
		switch v := i.(type) {
		case error:
			err = v
			break LOOP
		case *Resultset:
			if r == nil {
				r = v
			} else {
				if e := mergeResultset(r, v); err != nil {
					err = e
					break LOOP
				}
			}
		default:
			err = fmt.Errorf("invalid return type %T", i)
			break LOOP
		}
	}

	return r, err
}

func (c *conn) handleSelect(l *lex) error {
	var r *Resultset
	var err error
	switch strings.ToUpper(l.Get(1).Value) {
	case `LAST_INSERT_ID`:
		r, err = c.handleSelectLastInsertId(l)
	case `ROW_COUNT`:
		r, err = c.handleSelectRowCount(l)
	default:
		r, err = c.handleSelectDefault(l)
	}

	if err != nil {
		return err
	} else {
		return c.writeResultset(r)
	}
}
