package proxy

import (
	"bytes"
	"encoding/binary"
	"github.com/siddontang/golib/log"
	"net"
	"sync/atomic"
)

var DEFAULT_CAPABILITY uint32 = CLIENT_LONG_PASSWORD | CLIENT_LONG_FLAG |
	CLIENT_CONNECT_WITH_DB | CLIENT_PROTOCOL_41 |
	CLIENT_TRANSACTIONS | CLIENT_SECURE_CONNECTION

//client <-> proxy
type ClientConn struct {
	Conn

	server *Server

	connectionId uint32
	capability   uint32
	status       uint16
	charset      byte

	isAutoCommit  bool
	isTransaction bool

	user string
	db   string

	salt []byte

	msgs chan []byte

	quit    chan bool
	running bool

	nodeConns map[*DataNode]*ProxyConn
}

var BaseConnectionId uint32 = 10000

func NewClientConn(s *Server, c net.Conn) *ClientConn {
	conn := new(ClientConn)

	conn.server = s

	conn.conn = c
	conn.sequence = 0

	conn.connectionId = atomic.AddUint32(&BaseConnectionId, 1)

	conn.status = SERVER_STATUS_AUTOCOMMIT

	conn.isAutoCommit = true
	conn.isTransaction = false

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

		c.WritePacket(DumpError(err, c.capability))

		return err
	}

	if err := c.WritePacket(DumpOK(&OKPacket{0, 0, c.status, 0, ""}, c.capability)); err != nil {
		log.Error("write ok fail %s", err.Error())
		return err
	}

	c.sequence = 0

	return nil
}

func (c *ClientConn) writeInitialHandshake() error {
	data := make([]byte, 4, 128)

	//min version 10
	data = append(data, 10)

	//server version[00]
	data = append(data, ServerVersion...)
	data = append(data, 0)

	//connection id
	data = append(data, byte(c.connectionId), byte(c.connectionId>>8), byte(c.connectionId>>16), byte(c.connectionId>>24))

	//auth-plugin-data-part-1
	data = append(data, c.salt[0:8]...)

	//filter [00]
	data = append(data, 0)

	//capability flag lower 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))

	//charset, utf-8 default
	data = append(data, DEFAULT_UTF8_CHARSET)

	//status
	data = append(data, byte(c.status), byte(c.status>>8))

	//below 13 byte may not be used
	//capability flag upper 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY>>16), byte(DEFAULT_CAPABILITY>>24))

	//filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x15)

	//reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	//auth-plugin-data-part-2
	data = append(data, c.salt[8:]...)

	//filter [00]
	data = append(data, 0)

	return c.WritePacket(data)
}

func (c *ClientConn) readHandshakeResponse() error {
	data, err := c.ReadPacket()

	if err != nil {
		return err
	}

	pos := 0

	//capability
	c.capability = binary.LittleEndian.Uint32(data[:4])
	pos += 4

	//skip max packet size
	pos += 4

	//charset
	c.charset = data[pos]
	pos++

	//skip reserved 23[00]
	pos += 23

	//user name
	c.user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
	pos += len(c.user) + 1

	//auth length and auth
	authLen := int(data[pos])
	pos++
	auth := data[pos : pos+authLen]

	checkAuth := CalcPassword(c.salt, []byte(c.server.cfg.Password))

	if !bytes.Equal(auth, checkAuth) {
		return NewDefaultMySQLError(ER_ACCESS_DENIED_ERROR, c.RemoteAddr().String(), c.user)
	}

	pos += authLen

	if c.capability|CLIENT_CONNECT_WITH_DB > 0 {
		c.db = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(c.db) + 1

		//todo check db in schemas
		if c.server.schemas.GetSchema(c.db) == nil {
			return NewDefaultMySQLError(ER_BAD_DB_ERROR, c.db)
		}
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

func (c *ClientConn) handleReadPacket(data []byte) error {
	return nil
}

func (c *ClientConn) onRead() {
	for {
		data, err := c.ReadPacket()
		if err != nil {
			log.Error("read packet error %s", err.Error())
			return
		}

		if err := c.handleReadPacket(data); err != nil {

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
