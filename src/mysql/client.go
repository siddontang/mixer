package mysql

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/siddontang/golib/log"
	"io"
	"net"
)

var (
	ErrMalformPacket   = errors.New("write Malform error")
	ErrPayloadLength   = errors.New("invalid payload length")
	ErrPacketSequence  = errors.New("invalid packet sequence")
	ErrProtocolVersion = errors.New("invalid protocol version")
	ErrEOFPacket       = errors.New("eof packet")
	ErrNotSupported    = errors.New("not supported")
	ErrMismatchColumns = errors.New("column mismatch")
)

//proxy <-> mysql server
//refer go-sql-driver
type Client struct {
	address  string
	user     string
	password string
	db       string

	conn     net.Conn
	sequence uint8

	authData   []byte
	authName   string
	capability uint32
	charset    uint8
	status     uint16

	affectedRows uint64
	lastInsertId uint64
}

func NewClient() *Client {
	c := new(Client)

	return c
}

func (c *Client) Connect(address string, user string, password string, db string) error {
	c.address = address

	c.user = user
	c.password = password
	c.db = db

	c.conn = nil

	return c.ReConnect()
}

func (c *Client) ReConnect() error {
	if c.conn != nil {
		c.conn.Close()
	}

	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		log.Error("connect %s error %s", c.address, err.Error())
		return err
	}

	c.conn = conn
	c.sequence = 0

	c.authData = []byte{}
	c.authName = ""
	c.capability = 0
	c.charset = 0
	c.status = 0

	if err := c.readInitPacket(); err != nil {
		log.Error("read initial handshake packet error %s", err.Error())
		return err
	}

	if err := c.writeAuthPacket(); err != nil {
		log.Error("write auth response packet error %s", err.Error())
		return err
	}

	if err := c.readResultOKPacket(); err != nil {
		log.Error("read result ok packet error %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}

	c.conn = nil

	return nil
}

func (c *Client) readPacket() ([]byte, error) {
	header := make([]byte, 4)

	if _, err := io.ReadFull(c.conn, header); err != nil {
		return nil, err
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		log.Error("invalid payload length")
		return nil, ErrPayloadLength
	}

	sequence := uint8(header[3])

	if sequence != c.sequence {
		log.Error("invalid sequence %d != %d", sequence, c.sequence)
		return nil, ErrPacketSequence
	}

	c.sequence++

	data := make([]byte, length)
	if _, err := io.ReadFull(c.conn, data); err != nil {
		log.Error("read payload data error %s", err.Error())
		return nil, err
	} else {
		if length < MaxPayloadLen {
			return data, nil
		}

		var buf []byte
		buf, err = c.readPacket()
		if err != nil {
			log.Error("read packet error %s", err.Error())
			return nil, err
		} else {
			return append(data, buf...), nil
		}
	}
}

//data already have header
func (c *Client) writePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = c.sequence

		if n, err := c.conn.Write(data[:4+MaxPayloadLen]); err != nil {
			log.Error("write error %s", err.Error())
			return err
		} else if n != (4 + MaxPayloadLen) {
			log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
			return ErrMalformPacket
		} else {
			c.sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = c.sequence

	if n, err := c.conn.Write(data); err != nil {
		log.Error("write error %s", err.Error())
		return err
	} else if n != len(data) {
		log.Error("write error, write data number %d != %d", n, (4 + MaxPayloadLen))
		return ErrMalformPacket
	} else {
		c.sequence++
		return nil
	}
}

func (c *Client) handleOKPacket(data []byte) error {
	var n int

	// 0x00 [1 byte]

	// Affected rows [Length Coded Binary]
	c.affectedRows, _, n = readLengthEncodedInteger(data[1:])

	// Insert id [Length Coded Binary]
	c.lastInsertId, _, _ = readLengthEncodedInteger(data[1+n:])

	return nil
}

func (c *Client) handleErrorPacket(data []byte) error {
	if data[0] != ERR_Packet {
		return ErrMalformPacket
	}

	errorCode := binary.LittleEndian.Uint16(data[1:3])

	pos := 3
	//sql state marker and state
	//maker is #, state is 5 length
	if data[pos] == '#' {
		pos = 9
	}

	return &MySQLError{
		Code:    errorCode,
		Message: string(data[pos:]),
	}
}

func (c *Client) handleLocalInFilePacket(data []byte) error {
	//now we not support local in file protocol
	return ErrNotSupported
}

func (c *Client) readInitPacket() error {
	data, err := c.readPacket()
	if err != nil {
		return err
	}

	if data[0] == ERR_Packet {
		return c.handleErrorPacket(data)
	}

	if data[0] < MinProtocolVersion {
		log.Error("invalid protocol version %d < %d", data[0], MinProtocolVersion)
		return ErrProtocolVersion
	}

	//skip mysql version and connection id
	//mysql version end with 0x00
	//connection id length is 4
	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1 + 4

	c.authData = append(c.authData, data[pos:pos+8]...)

	//skip filter
	pos += 8 + 1

	//capability lower 2 bytes
	c.capability = uint32(binary.LittleEndian.Uint16(data[pos : pos+2]))

	pos += 2

	if len(data) > pos {
		c.charset = data[pos]
		pos += 1
		c.status = binary.LittleEndian.Uint16(data[pos : pos+2])

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
		c.authData = append(c.authData, data[pos:pos+12]...)
	}

	return nil
}

func (c *Client) writeAuthPacket() error {
	// Adjust client capability flags based on server support
	capability := uint32(
		CLIENT_PROTOCOL_41 |
			CLIENT_SECURE_CONNECTION |
			CLIENT_LONG_PASSWORD |
			CLIENT_TRANSACTIONS |
			CLIENT_LOCAL_FILES,
	)

	if (c.capability & CLIENT_LONG_FLAG) != 0 {
		capability |= CLIENT_LONG_FLAG
	}

	//packet length
	//capbility 4
	//max-packet size 4
	//charset 1
	//reserved all[0] 23
	length := 4 + 4 + 1 + 23

	//username
	length += len(c.user) + 1

	//we only support secure connection
	auth := calcPassword(c.authData, []byte(c.password))

	length += 1 + len(auth)

	if len(c.db) > 0 {
		capability |= CLIENT_CONNECT_WITH_DB

		length += len(c.db) + 1
	}

	data := make([]byte, length+4)

	// Add the packet header  [24bit length + 1 byte sequence]
	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = c.sequence

	// capability [32 bit]
	data[4] = byte(capability)
	data[5] = byte(capability >> 8)
	data[6] = byte(capability >> 16)
	data[7] = byte(capability >> 24)

	// MaxPacketSize [32 bit] (none)
	//data[8] = 0x00
	//data[9] = 0x00
	//data[10] = 0x00
	//data[11] = 0x00

	// Charset [1 byte]
	data[12] = c.charset

	// Filler [23 bytes] (all 0x00)
	pos := 13 + 23

	// User [null terminated string]
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

	// Send Auth packet
	return c.writePacket(data)
}

func (c *Client) readResultOKPacket() error {
	data, err := c.readPacket()
	if err == nil {
		switch data[0] {
		case OK_Packet:
			return c.handleOKPacket(data)
		case EOF_Packet:
			return ErrEOFPacket
		case ERR_Packet:
			return c.handleErrorPacket(data)
		}
	}
	return err
}

func (c *Client) readTextResultSetColumns(count uint64) (columns [][]byte, err error) {
	var i uint64
	var data []byte

	columns = make([][]byte, count)

	for {
		data, err = c.readPacket()
		if err != nil {
			return
		}

		// EOF Packet
		if data[0] == EOF_Packet && len(data) < 9 {
			if i != count {
				log.Error("ColumnsCount mismatch n:%d len:%d", count, len(columns))
				err = ErrMismatchColumns
			}
			return
		}

		columns[i] = data

		i++
	}
}

func (c *Client) readTextResultSetRows() (rows [][]byte, err error) {
	var data []byte

	rows = make([][]byte, 0)

	for {
		data, err = c.readPacket()

		if err != nil {
			return
		}

		// EOF Packet
		if data[0] == EOF_Packet && len(data) < 9 {
			return
		}

		rows = append(rows, data)
	}
}

func (c *Client) readTextResultSetPacket() (columns [][]byte, rows [][]byte, err error) {
	var data []byte

	data, err = c.readPacket()
	if err != nil {
		return
	}

	switch data[0] {
	case OK_Packet:
		err = c.handleOKPacket(data)
		return
	case ERR_Packet:
		err = c.handleErrorPacket(data)
		return
	case LocalInFile_Packet:
		err = c.handleLocalInFilePacket(data[1:])
		return
	}

	// column count
	count, _, n := readLengthEncodedInteger(data)
	if n-len(data) != 0 {
		err = ErrMalformPacket
		return
	}

	columns, err = c.readTextResultSetColumns(count)
	if err != nil {
		return
	}

	rows, err = c.readTextResultSetRows()

	return
}

func (c *Client) writeCommandPacket(command byte) error {
	c.sequence = 0

	return c.writePacket([]byte{
		0x01, //1 bytes long
		0x00,
		0x00,
		0x00, //sequence
		command,
	})
}

func (c *Client) writeCommandPacketStr(command byte, arg string) error {
	c.sequence = 0

	length := len(arg) + 1

	data := make([]byte, length+4)

	//header, will be calculated in writePacket
	//data[0] = byte(length)
	//data[1] = byte(length >> 8)
	//data[2] = byte(length >> 16)
	//data[3] = c.sequence

	data[4] = command

	copy(data[5:], arg)

	return c.writePacket(data)
}

func (c *Client) writeCommandPacketUint32(command byte, arg uint32) error {
	c.sequence = 0

	return c.writePacket([]byte{
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
