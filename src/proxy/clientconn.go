package proxy

import (
	"bytes"
	"encoding/binary"
	"net"
	"sync/atomic"
)

type ClientConn struct {
	Conn

	server *Server

	ConnectionId uint32
	Capability   uint32
	Charset      byte

	User string
	DB   string

	salt []byte
}

var BaseConnectionId = 1000

func NewClientConn(s *Server, c net.Conn) *ClientConn {
	conn := new(ClientConn)

	conn.server = s

	conn.conn = c
	conn.sequence = 0

	conn.ConnectionId = atomic.AddUint32(&BaseConnectionId, 1)

	conn.salt = RandomBuf(20)

	return conn
}

func (c *ClientConn) Handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		log.Error("send initial handshake error %s", err.Error())
		return
	}

	if err := c.readHandshakeResponse(); err != nil {
		log.Error("recv handshake response error %s", err.Error())
		return
	}

	if err := c.WriteOK(); err != nil {
		log.Error("write ok error %s", err.Error())
		return
	}

	return nil
}

func (c *ClientConn) writeInitialHandshake() error {
	buf := make([]byte, 0, 128)

	buf = append(buf, 0, 0, 0, 0)

	//min version 10
	buf = append(buf, 10)

	//server version[00]
	buf = append(buf, ServerVersion...)
	buf = append(buf, 0)

	//connection id
	buf = append(buf, byte(c.ConnectionId), byte(c.ConnectionId>>8), byte(c.ConnectionId>>16), byte(c.ConnectionId>>24))

	//auth-plugin-data-part-1
	buf = append(buf, c.salt[0:8]...)

	//filter [00]
	buf = append(buf, 0)

	//capability flag lower 2 bytes, using default capability here
	buf = append(buf, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))

	//charset, utf-8 default
	buf = append(buf, DEFAULT_CAPABILITY)

	//status, only auto commit here
	buf = append(buf, SERVER_STATUS_AUTOCOMMIT)

	//below 13 byte may not be used
	//capability flag upper 2 bytes, using default capability here
	buf = append(buf, byte(DEFAULT_CAPABILITY>>16), byte(DEFAULT_CAPABILITY>>24))

	//filter [15], for wireshark dump, value is 15
	buf = append(buf, 15)

	//reserved 10 [00]
	buf = append(buf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	//auth-plugin-data-part-2
	buf = append(buf, c.salt[8:]...)

	//filter [00]
	buf = append(buf, 0)

	//auth name
	buf = append(buf, AUTH_NAME...)
	buf = append(buf, 0)

	return c.WritePacket(buf)
}

func (c *ClientConn) readHandshakeResponse() error {
	buf, err := c.ReadPacket()

	if err != nil {
		return err
	}

	pos := 0

	//capability
	c.Capability = binary.LittleEndian.Uint32(buf[:4])
	pos += 4

	//skip max packet size
	pos += 4

	//charset
	c.Charset = buf[pos]
	pos++

	//skip reserved 23[00]
	pos += 23

	//user name
	c.User = string(buf[pos:bytes.IndexByte(buf[pos:], 0)])
	pos += len(c.User) + 1

	//auth length and auth
	authLen := int(buf[pos])
	pos++
	auth := buf[pos : pos+authLen]

	checkAuth := calcPassword(c.salt, []byte(c.server.cfg.Password))

	if !bytes.Equal(auth, checkAuth) {

	}

	pos += authLen

	if c.Capability | CLIENT_CONNECT_WITH_DB {
		c.DB = string(buf[pos:bytes.IndexByte(buf[pos:], 0)])
		pos += len(c.DB) + 1
	}

	//auth name
	authName := string(buf[pos:bytes.IndexByte(buf[pos:], 0)])

	return nil
}

func (c *ClientConn) WriteOK() error {

}
