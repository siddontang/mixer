package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/siddontang/go-log/log"
	. "github.com/siddontang/mixer/mysql"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
)

var DEFAULT_CAPABILITY uint32 = CLIENT_LONG_PASSWORD | CLIENT_LONG_FLAG |
	CLIENT_CONNECT_WITH_DB | CLIENT_PROTOCOL_41 |
	CLIENT_TRANSACTIONS | CLIENT_SECURE_CONNECTION

//client <-> proxy
type conn struct {
	pkg *PacketIO

	c net.Conn

	sync.Mutex

	server *Server

	capability uint32

	connectionId uint32

	status    uint16
	collation CollationId
	charset   string

	user string
	db   string

	salt []byte

	curSchema *schema

	txConns map[*node]*SqlConn

	stmtId uint32
	stmts  map[uint32]*stmt

	closed bool

	lastInsertId int64
	affectedRows int64
}

var baseConnId uint32 = 10000

func newconn(s *Server, co net.Conn) *conn {
	c := new(conn)

	c.c = co

	c.pkg = NewPacketIO(co)

	c.server = s

	c.c = co
	c.pkg.Sequence = 0

	c.connectionId = atomic.AddUint32(&baseConnId, 1)

	c.status = SERVER_STATUS_AUTOCOMMIT

	c.salt, _ = RandomBuf(20)

	c.txConns = make(map[*node]*SqlConn)
	c.stmts = make(map[uint32]*stmt)

	c.closed = false

	c.stmtId = 0

	c.collation = DEFAULT_COLLATION_ID
	c.charset = DEFAULT_CHARSET

	return c
}

func (c *conn) Handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		log.Error("send initial handshake error %s", err.Error())
		return err
	}

	if err := c.readHandshakeResponse(); err != nil {
		log.Error("recv handshake response error %s", err.Error())

		c.writeError(err)

		return err
	}

	if err := c.writeOK(nil); err != nil {
		log.Error("write ok fail %s", err.Error())
		return err
	}

	c.pkg.Sequence = 0

	return nil
}

func (c *conn) Close() error {
	if c.closed {
		return nil
	}

	c.c.Close()

	c.rollback()

	c.closed = true

	return nil
}

func (c *conn) writeInitialHandshake() error {
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
	data = append(data, uint8(DEFAULT_COLLATION_ID))

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

	return c.writePacket(data)
}

func (c *conn) readPacket() ([]byte, error) {
	return c.pkg.ReadPacket()
}

func (c *conn) writePacket(data []byte) error {
	return c.pkg.WritePacket(data)
}

func (c *conn) readHandshakeResponse() error {
	data, err := c.readPacket()

	if err != nil {
		return err
	}

	pos := 0

	//capability
	c.capability = binary.LittleEndian.Uint32(data[:4])
	pos += 4

	//skip max packet size
	pos += 4

	//charset, skip, if you want to use another charset, use set names
	//c.collation = CollationId(data[pos])
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
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, c.c.RemoteAddr().String(), c.user)
	}

	pos += authLen

	if c.capability|CLIENT_CONNECT_WITH_DB > 0 {
		db := string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(c.db) + 1

		if err := c.useDB(db); err != nil {
			return err
		}
	}

	return nil
}

func (c *conn) Run() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			log.Error("%v, %s", err, buf)
		}

		c.Close()
	}()

	for {
		data, err := c.readPacket()

		if err != nil {
			return
		}

		if err := c.dispatch(data); err != nil {
			log.Error("dispatch error %s", err.Error())
			if err != ErrBadConn {
				c.writeError(err)
			}
		}

		c.pkg.Sequence = 0
	}
}

func (c *conn) dispatch(data []byte) error {
	cmd := data[0]
	data = data[1:]

	switch cmd {
	case COM_QUERY:
		return c.handleQuery(data)
	case COM_PING:
		return c.writeOK(nil)
	case COM_INIT_DB:
		return c.useDB(string(data))
	case COM_STMT_PREPARE:
		return c.handleStmtPrepare(data)
	case COM_STMT_EXECUTE:
		return c.handleStmtExecute(data)
	case COM_STMT_CLOSE:
		return c.handleStmtClose(data)
	case COM_STMT_SEND_LONG_DATA:
		return c.handleStmtSendLongData(data)
	case COM_STMT_RESET:
		return c.handleStmtReset(data)
	default:
		msg := fmt.Sprintf("command %d not supported now", cmd)
		return NewError(ER_UNKNOWN_ERROR, msg)
	}

	return nil
}

func (c *conn) useDB(db string) error {
	if s := c.server.schemas.GetSchema(db); s == nil {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	} else {
		c.curSchema = s
		c.db = db
	}
	return nil
}

func (c *conn) writeOK(r *Result) error {
	if r == nil {
		r = &Result{Status: c.status}
	}
	data := make([]byte, 4, 32)

	data = append(data, OK_HEADER)

	data = append(data, PutLengthEncodedInt(r.AffectedRows)...)
	data = append(data, PutLengthEncodedInt(r.InsertId)...)

	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status)>>8)
		data = append(data, 0, 0)
	}

	return c.writePacket(data)
}

func (c *conn) writeError(e error) error {
	var m *SqlError
	var ok bool
	if m, ok = e.(*SqlError); !ok {
		m = NewError(ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}

	data = append(data, m.Message...)

	return c.writePacket(data)
}

func (c *conn) writeEOF(status uint16) error {
	data := make([]byte, 4, 9)

	data = append(data, EOF_HEADER)
	if c.capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status)>>8)
	}

	return c.writePacket(data)
}
