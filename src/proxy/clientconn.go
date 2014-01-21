package proxy

import (
	"bytes"
	"encoding/binary"
	"github.com/siddontang/golib/log"
	"net"
	"sync/atomic"
)

var DEFAULT_CAPABILITY uint32 = CLIENT_LONG_PASSWORD | CLIENT_FOUND_ROWS | CLIENT_LONG_FLAG |
	CLIENT_CONNECT_WITH_DB | CLIENT_PROTOCOL_41 | CLIENT_TRANSACTIONS | CLIENT_SECURE_CONNECTION

//client <-> proxy
type ClientConn struct {
	Conn

	server *Server

	connectionId uint32
	capability   uint32
	status       uint16
	charset      byte

	user string
	db   string

	salt []byte

	msgs chan []byte

	quit    chan bool
	running bool
}

var BaseConnectionId uint32 = 1000

func NewClientConn(s *Server, c net.Conn) *ClientConn {
	conn := new(ClientConn)

	conn.server = s

	conn.conn = c
	conn.sequence = 0

	conn.connectionId = atomic.AddUint32(&BaseConnectionId, 1)

	conn.status = SERVER_STATUS_AUTOCOMMIT

	conn.salt, _ = RandomBuf(20)

	conn.quit = make(chan bool)

	conn.msgs = make(chan []byte, 4)

	return conn
}

func (c *ClientConn) Handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		log.Error("send initial handshake error %s", err.Error())
		return err
	}

	if err := c.readHandshakeResponse(); err != nil {
		log.Error("recv handshake response error %s", err.Error())

		c.WritePacket(c.BuildError(err, c.capability))

		return err
	}

	if err := c.WritePacket(c.BuildOK(0, 0, c.capability, c.status, 0, "")); err != nil {
		log.Error("write ok fail %s", err.Error())
		return err
	}

	return nil
}

func (c *ClientConn) writeInitialHandshake() error {
	buf := make([]byte, 4, 128)

	//min version 10
	buf = append(buf, 10)

	//server version[00]
	buf = append(buf, ServerVersion...)
	buf = append(buf, 0)

	//connection id
	buf = append(buf, byte(c.connectionId), byte(c.connectionId>>8), byte(c.connectionId>>16), byte(c.connectionId>>24))

	//auth-plugin-data-part-1
	buf = append(buf, c.salt[0:8]...)

	//filter [00]
	buf = append(buf, 0)

	//capability flag lower 2 bytes, using default capability here
	buf = append(buf, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))

	//charset, utf-8 default
	buf = append(buf, c.charset)

	//status
	buf = append(buf, byte(c.status), byte(c.status>>8))

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
	c.capability = binary.LittleEndian.Uint32(buf[:4])
	pos += 4

	//skip max packet size
	pos += 4

	//charset
	c.charset = buf[pos]
	pos++

	//skip reserved 23[00]
	pos += 23

	//user name
	c.user = string(buf[pos:bytes.IndexByte(buf[pos:], 0)])
	pos += len(c.user) + 1

	//auth length and auth
	authLen := int(buf[pos])
	pos++
	auth := buf[pos : pos+authLen]

	checkAuth := CalcPassword(c.salt, []byte(c.server.cfg.Password))

	if !bytes.Equal(auth, checkAuth) {
		return NewDefaultMySQLError(ER_ACCESS_DENIED_ERROR, c.RemoteAddr().String(), c.user)
	}

	pos += authLen

	if c.capability|CLIENT_CONNECT_WITH_DB > 0 {
		c.db = string(buf[pos:bytes.IndexByte(buf[pos:], 0)])
		pos += len(c.db) + 1

		if c.server.nodes.GetNode(c.db) == nil {
			return NewDefaultMySQLError(ER_BAD_DB_ERROR, c.db)
		}
	}

	//auth name
	authName := string(buf[pos:bytes.IndexByte(buf[pos:], 0)])
	if authName != AUTH_NAME {
		return NewDefaultMySQLError(ER_ACCESS_DENIED_ERROR, c.RemoteAddr().String(), c.user)
	}

	return nil
}

func (c *ClientConn) Run() {
	c.running = true

	go c.onWrite()

	c.onRead()

	c.running = false

	close(c.quit)
}

func (c *ClientConn) handleReadPacket(buf []byte) error {
	return nil
}

func (c *ClientConn) onRead() {
	for {
		buf, err := c.ReadPacket()
		if err != nil {
			log.Error("read packet error %s", err.Error())
			return
		}

		if err := c.handleReadPacket(buf); err != nil {

		}
	}
}

func (c *ClientConn) onWrite() {
	for {
		select {
		case msg := <-c.msgs:
			if err := c.WritePacket(msg); err != nil {
				log.Error("write packet error %s", err.Error())
				return
			}
		case <-c.quit:
			return
		}
	}
}
