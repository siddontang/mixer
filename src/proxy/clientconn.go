package proxy

import (
	"net"
)

type ClientConn struct {
	Conn

	server *Server
}

func NewClientConn(s *Server, c net.Conn) *ClientConn {
	conn := new(ClientConn)

	conn.server = s

	conn.conn = c
	conn.sequence = 0

	return conn
}

func (conn *ClientConn) Handshake() error {
	if err := conn.writeInitialHandshake(); err != nil {
		log.Error("send initial handshake error %s", err.Error())
		return
	}

	if err := conn.readHandshakeResponse(); err != nil {
		log.Error("recv handshake response error %s", err.Error())
		return
	}

	if err := conn.WriteOK(); err != nil {
		log.Error("write ok error %s", err.Error())
		return
	}

}

func (c *ClientConn) writeInitialHandshake() error {
	return nil
}

func (c *ClientConn) readHandshakeResponse() error {
	return nil
}

func (c *ClientConn) WriteOK() error {

}
