package proxy

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/siddontang/golib/log"
	"io"
	"mysql"
	"net"
	"sync/atomic"
)

var DEFAULT_CAPABILITY uint32 = mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_LONG_FLAG |
	mysql.CLIENT_CONNECT_WITH_DB | mysql.CLIENT_PROTOCOL_41 |
	mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_SECURE_CONNECTION

//client <-> proxy
type ClientConn struct {
	mysql.Conn

	server *Server

	connectionId uint32

	status  uint16
	charset byte

	user string
	db   string

	salt []byte

	schema *Schema

	nodeConns map[*DataNode]*mysql.Client
}

var BaseConnectionId uint32 = 10000

func NewClientConn(s *Server, c net.Conn) *ClientConn {
	conn := new(ClientConn)

	conn.server = s

	conn.NetConn = c
	conn.Sequence = 0

	conn.connectionId = atomic.AddUint32(&BaseConnectionId, 1)

	conn.status = mysql.SERVER_STATUS_AUTOCOMMIT

	conn.salt, _ = mysql.RandomBuf(20)

	conn.nodeConns = make(map[*DataNode]*mysql.Client)

	return conn
}

func (c *ClientConn) Handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		log.Error("send initial handshake error %s", err.Error())
		return err
	}

	if err := c.readHandshakeResponse(); err != nil {
		log.Error("recv handshake response error %s", err.Error())

		c.WriteError(err)

		return err
	}

	if err := c.WriteOK(&mysql.OKPacket{0, 0, c.status, 0, ""}); err != nil {
		log.Error("write ok fail %s", err.Error())
		return err
	}

	c.Sequence = 0

	return nil
}

func (c *ClientConn) Close() error {
	c.NetConn.Close()

	//connection closed but proxy connection may be in trans, cancel
	for node, conn := range c.nodeConns {
		if _, err := conn.Rollback(); err != nil {
			log.Error("node %s rollback error %s", node.name, err.Error())
		}
	}

	c.clearNodeConns()
	return nil
}

func (c *ClientConn) clearNodeConns() {
	for n, v := range c.nodeConns {
		n.PushConn(v)
		delete(c.nodeConns, n)
	}
}

func (c *ClientConn) writeInitialHandshake() error {
	data := make([]byte, 4, 128)

	//min version 10
	data = append(data, 10)

	//server version[00]
	data = append(data, mysql.ServerVersion...)
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
	data = append(data, mysql.DEFAULT_UTF8_CHARSET)

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
	c.Capability = binary.LittleEndian.Uint32(data[:4])
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

	checkAuth := mysql.CalcPassword(c.salt, []byte(c.server.cfg.Password))

	if !bytes.Equal(auth, checkAuth) {
		return mysql.NewDefaultError(mysql.ER_ACCESS_DENIED_ERROR, c.RemoteAddr().String(), c.user)
	}

	pos += authLen

	if c.Capability|mysql.CLIENT_CONNECT_WITH_DB > 0 {
		db := string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(c.db) + 1

		if err := c.useDB(db); err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientConn) Run() {
	for {
		data, err := c.ReadPacket()

		if err != nil {
			if err != io.EOF {
				log.Error("read packet error %s, close", err.Error())
			}
			return
		}

		if err := c.dispatch(data); err != nil {
			log.Error("dispatch error %s", err.Error())
			c.WriteError(err)
		}

		c.Sequence = 0
	}
}

func (c *ClientConn) dispatch(data []byte) error {
	switch data[0] {
	case mysql.COM_QUERY:
		return c.handleQuery(data[1:])
	case mysql.COM_PING:
		c.WriteOK(&mysql.OKPacket{Status: c.status})
		return nil
	case mysql.COM_INIT_DB:
		return c.useDB(string(data[1:]))
	default:
		msg := fmt.Sprintf("command %d not supported now", data[0])
		return mysql.NewError(mysql.ER_UNKNOWN_ERROR, msg)
	}

	return nil
}

func (c *ClientConn) useDB(db string) error {
	if s := c.server.schemas.GetSchema(db); s == nil {
		return mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db)
	} else {
		c.schema = s
		c.db = db
	}
	return nil
}
