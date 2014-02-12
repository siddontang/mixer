package mysql

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"time"
)

var (
	pingPeriod = int64(time.Second * 30)
)

//proxy <-> mysql server
type conn struct {
	PacketIO

	addr     string
	user     string
	password string
	db       string

	capability uint32

	//status  uint16
	charset byte
	salt    []byte

	lastPing int64
}

func (c *conn) Connect(addr string, user string, password string, db string) error {
	c.addr = addr
	c.user = user
	c.password = password
	c.db = db

	//use utf8
	c.charset = DEFAULT_UTF8_CHARSET

	return c.ReConnect()
}

func (c *conn) ReConnect() error {
	if c.Conn != nil {
		c.Conn.Close()
	}

	netConn, err := net.Dial("tcp", c.addr)
	if err != nil {
		errLog("connect %s error %s", c.addr, err.Error())
		return err
	}

	c.Conn = netConn
	c.Sequence = 0

	if err := c.readInitialHandshake(); err != nil {
		errLog("read initial handshake error %s", err.Error())
		c.Conn.Close()
		return err
	}

	if err := c.writeAuthHandshake(); err != nil {
		errLog("write auth handshake error %s", err.Error())
		c.Conn.Close()

		return err
	}

	if _, err := c.ReadOK(); err != nil {
		errLog("read ok error %s", err.Error())
		c.Conn.Close()

		return err
	}

	c.lastPing = time.Now().Unix()

	return nil
}

func (c *conn) Close() error {
	if c.Conn != nil {
		c.Conn.Close()
	}

	return nil
}

func (c *conn) readInitialHandshake() error {
	data, err := c.ReadPacket()
	if err != nil {
		return err
	}

	if data[0] == ERR_HEADER {
		return errors.New("read initial handshake error")
	}

	if data[0] < MinProtocolVersion {
		err := fmt.Errorf("invalid protocol version %d, must >= 10", data[0])
		errLog(err.Error())
		return err
	}

	//skip mysql version and connection id
	//mysql version end with 0x00
	//connection id length is 4
	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1 + 4

	c.salt = append(c.salt, data[pos:pos+8]...)

	//skip filter
	pos += 8 + 1

	//capability lower 2 bytes
	c.capability = uint32(binary.LittleEndian.Uint16(data[pos : pos+2]))

	pos += 2

	if len(data) > pos {
		//skip server charset
		//c.charset = data[pos]
		pos += 1

		//c.status = binary.LittleEndian.Uint16(data[pos : pos+2])
		pos += 2

		c.capability = uint32(binary.LittleEndian.Uint16(data[pos:pos+2]))<<16 | c.capability

		pos += 2

		//skip auth data len or [00]
		//skip reserved (all [00])
		pos += 10 + 1

		// The documentation is ambiguous about the length.
		// The official Python library uses the fixed length 12
		// mysql-proxy also use 12
		// which is not documented but seems to work.
		c.salt = append(c.salt, data[pos:pos+12]...)
	}

	return nil
}

func (c *conn) writeAuthHandshake() error {
	// Adjust client capability flags based on server support
	capability := CLIENT_PROTOCOL_41 | CLIENT_SECURE_CONNECTION |
		CLIENT_LONG_PASSWORD | CLIENT_TRANSACTIONS | CLIENT_LONG_FLAG

	capability &= c.capability

	//packet length
	//capbility 4
	//max-packet size 4
	//charset 1
	//reserved all[0] 23
	length := 4 + 4 + 1 + 23

	//username
	length += len(c.user) + 1

	//we only support secure connection
	auth := CalcPassword(c.salt, []byte(c.password))

	length += 1 + len(auth)

	if len(c.db) > 0 {
		capability |= CLIENT_CONNECT_WITH_DB

		length += len(c.db) + 1
	}

	c.capability = capability

	data := make([]byte, length+4)

	//capability [32 bit]
	data[4] = byte(capability)
	data[5] = byte(capability >> 8)
	data[6] = byte(capability >> 16)
	data[7] = byte(capability >> 24)

	//MaxPacketSize [32 bit] (none)
	//data[8] = 0x00
	//data[9] = 0x00
	//data[10] = 0x00
	//data[11] = 0x00

	//Charset [1 byte]
	data[12] = c.charset

	//Filler [23 bytes] (all 0x00)
	pos := 13 + 23

	//User [null terminated string]
	if len(c.user) > 0 {
		pos += copy(data[pos:], c.user)
	}
	//data[pos] = 0x00
	pos++

	// auth [length encoded integer]
	data[pos] = byte(len(auth))
	pos += 1 + copy(data[pos+1:], auth)

	// db [null terminated string]
	if len(c.db) > 0 {
		pos += copy(data[pos:], c.db)
		//data[pos] = 0x00
	}

	return c.WritePacket(data)
}

func (c *conn) WriteCommand(command byte) error {
	c.Sequence = 0

	return c.WritePacket([]byte{
		0x01, //1 bytes long
		0x00,
		0x00,
		0x00, //sequence
		command,
	})
}

func (c *conn) WriteCommandBuf(command byte, arg []byte) error {
	c.Sequence = 0

	length := len(arg) + 1

	data := make([]byte, length+4)

	data[4] = command

	copy(data[5:], arg)

	return c.WritePacket(data)
}

func (c *conn) WriteCommandStr(command byte, arg string) error {
	c.Sequence = 0

	length := len(arg) + 1

	data := make([]byte, length+4)

	data[4] = command

	copy(data[5:], arg)

	return c.WritePacket(data)
}

func (c *conn) WriteCommandUint32(command byte, arg uint32) error {
	c.Sequence = 0

	return c.WritePacket([]byte{
		0x05, //5 bytes long
		0x00,
		0x00,
		0x00, //sequence

		command,

		byte(arg),
		byte(arg >> 8),
		byte(arg >> 16),
		byte(arg >> 24),
	})
}

func (c *conn) WriteCommandStrStr(command byte, arg1 string, arg2 string) error {
	c.Sequence = 0

	data := make([]byte, 4, 6+len(arg1)+len(arg2))

	data = append(data, command)
	data = append(data, arg1...)
	data = append(data, 0)
	data = append(data, arg2...)

	return c.WritePacket(data)
}

func (c *conn) Ping() error {
	n := time.Now().Unix()

	if n-c.lastPing > pingPeriod {
		if err := c.WriteCommand(COM_PING); err != nil {
			return err
		}

		if _, err := c.ReadOK(); err != nil {
			return err
		}
	}

	c.lastPing = n

	return nil
}

func (c *conn) Exec(command string, args ...interface{}) (*Result, error) {
	if len(args) == 0 {
		if p, err := c.exec(command); err != nil {
			return nil, err
		} else {
			//todo, for strict_mode treat warning as error
			return &Result{Status: p.Status, InsertId: p.LastInsertId,
				AffectedRows: p.AffectedRows}, nil
		}
	} else {
		if s, err := c.Prepare(command); err != nil {
			return nil, err
		} else {
			var r *Result
			r, err = s.Exec(args...)
			s.Close()
			return r, err
		}
	}
}

func (c *conn) exec(query string) (*OKPacket, error) {
	if err := c.WriteCommandStr(COM_QUERY, query); err != nil {
		return nil, err
	}

	return c.ReadOK()
}

func (c *conn) Begin() error {
	_, err := c.exec("begin")
	return err
}

func (c *conn) Commit() error {
	_, err := c.exec("commit")
	return err
}

func (c *conn) Rollback() error {
	_, err := c.exec("rollback")
	return err
}

func (c *conn) FieldList(table, fieldWildcard string) ([]Field, error) {
	if err := c.WriteCommandStrStr(COM_FIELD_LIST, table, fieldWildcard); err != nil {
		return nil, err
	}

	data, err := c.ReadPacket()
	if err != nil {
		return nil, err
	}

	columns := make([]FieldPacket, 0)

	if data[0] == ERR_HEADER {
		return nil, LoadError(data)
	} else if data[0] == EOF_HEADER && len(data) <= 5 {
		return []Field{}, nil
	}

	columns = append(columns, data)

	for {
		data, err = c.ReadPacket()
		if err != nil {
			return nil, err
		}

		// EOF Packet
		if data[0] == EOF_HEADER && len(data) <= 5 {
			break
		}

		columns = append(columns, data)
	}

	f := make([]Field, len(columns))

	for i := range columns {
		f[i], err = columns[i].Parse()
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func (c *conn) Query(command string, args ...interface{}) (*Resultset, error) {
	if len(args) == 0 {
		return c.query(command)
	} else {
		if s, err := c.Prepare(command); err != nil {
			return nil, err
		} else {
			var r *Resultset
			r, err = s.Query(args...)
			s.Close()
			return r, err
		}
	}
}

func (c *conn) query(query string) (*Resultset, error) {
	if err := c.WriteCommandStr(COM_QUERY, query); err != nil {
		return nil, err
	}

	if r, err := c.readResultset(); err != nil {
		return nil, err
	} else {
		return r.Parse(false)
	}
}

func (c *conn) readResultset() (*ResultsetPacket, error) {
	data, err := c.ReadPacket()
	if err != nil {
		return nil, err
	}

	result := new(ResultsetPacket)

	switch data[0] {
	case OK_HEADER:
		return result, nil
	case ERR_HEADER:
		return nil, LoadError(data)
	case LocalInFile_HEADER:
		return nil, ErrMalformPacket
	}

	// column count
	count, _, n := LengthEncodedInt(data)

	if n-len(data) != 0 {
		return nil, ErrMalformPacket
	}

	result.Fields = make([]FieldPacket, count)
	result.Rows = make([]RowPacket, 0)

	if err := c.readResultColumns(result); err != nil {
		return nil, err
	}

	if err := c.readResultRows(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *conn) readResultColumns(result *ResultsetPacket) (err error) {
	var i int = 0
	var data []byte

	for {
		data, err = c.ReadPacket()
		if err != nil {
			return
		}

		// EOF Packet
		if data[0] == EOF_HEADER && len(data) <= 5 {
			if i != len(result.Fields) {
				errLog("ColumnsCount mismatch n:%d len:%d", i, len(result.Fields))
				err = ErrMalformPacket
			}

			return
		}

		result.Fields[i] = data

		i++
	}
}

func (c *conn) readResultRows(result *ResultsetPacket) (err error) {
	var data []byte

	for {
		data, err = c.ReadPacket()

		if err != nil {
			return
		}

		// EOF Packet
		if data[0] == EOF_HEADER && len(data) <= 5 {
			if c.capability&CLIENT_PROTOCOL_41 > 0 {
				//result.Warnings = binary.LittleEndian.Uint16(data[1:])
				//todo add strict_mode, warning will be treat as error
				result.Status = binary.LittleEndian.Uint16(data[3:])
			}

			return
		}

		result.Rows = append(result.Rows, data)
	}
}

func (c *conn) readUntilEOF() (err error) {
	var data []byte

	for {
		data, err = c.ReadPacket()

		if err != nil {
			return
		}

		// EOF Packet
		if data[0] == EOF_HEADER && len(data) <= 5 {
			return
		}
	}
	return
}
