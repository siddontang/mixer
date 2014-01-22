package proxy

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/siddontang/golib/log"
	"net"
)

var (
	ErrProtocolVersion = errors.New("invalid protocol version, only support >= 10")
)

//proxy <-> mysql server
type ProxyConn struct {
	Conn

	server *Server

	addr     string
	user     string
	password string
	db       string

	capability uint32
	status     uint16
	charset    byte
	salt       []byte
}

func NewProxyConn(s *Server) *ProxyConn {
	c := new(ProxyConn)

	c.server = s

	return c
}

func (c *ProxyConn) Connect(addr string, user string, password string, db string) error {
	c.addr = addr
	c.user = user
	c.password = password
	c.db = db

	return c.ReConnect()
}

func (c *ProxyConn) ReConnect() error {
	if c.conn != nil {
		c.conn.Close()
	}

	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		log.Error("connect %s error %s", c.addr, err.Error())
		return err
	}

	c.conn = conn
	c.sequence = 0

	if err := c.readInitialHandshake(); err != nil {
		log.Error("read initial handshake error %s", err.Error())
		return err
	}

	if err := c.writeAuthHandshake(); err != nil {
		log.Error("write auth handshake error %s", err.Error())
		return err
	}

	if _, err := c.readOK(); err != nil {
		log.Error("read ok error %s", err.Error())
		return err
	}

	return nil
}

func (c *ProxyConn) readInitialHandshake() error {
	data, err := c.ReadPacket()
	if err != nil {
		return err
	}

	if data[0] == ERR_HEADER {
		return errors.New("read initial handshake error")
	}

	if data[0] < MinProtocolVersion {
		log.Error("invalid protocol version %d", data[0])
		return ErrProtocolVersion
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
		c.salt = append(c.salt, data[pos:pos+12]...)
	}

	return nil
}

func (c *ProxyConn) writeAuthHandshake() error {
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

	return c.WritePacket(data)
}

func (c *ProxyConn) readOK() (*OKPacket, error) {
	data, err := c.ReadPacket()
	if err != nil {
		return nil, err
	}

	switch data[0] {
	case OK_HEADER:
		return LoadOK(data, c.capability), nil
	case ERR_HEADER:
		return nil, LoadError(data, c.capability)
	default:
		return nil, ErrInvalidOKPacket
	}
}
