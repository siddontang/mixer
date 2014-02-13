package proxy

import (
	"lib/log"
	"net"
	"runtime"
)

type Server struct {
	cfg *Config

	addr     string
	user     string
	password string

	nodes   DataNodes
	schemas Schemas

	running bool

	listener net.Listener
}

func NewServer(cfg *Config) *Server {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.ConfigServer.Addr
	s.user = cfg.ConfigServer.User
	s.password = cfg.ConfigServer.Password

	s.nodes = NewDataNodes(s)
	s.schemas = NewSchemas(s, s.nodes)

	return s
}

func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Error("listen error %s", err.Error())
		return err
	}

	s.running = true

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Error("accept error %s", err.Error())
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *Server) Stop() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) onConn(c net.Conn) {
	conn := NewClientConn(s, c)

	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Error("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}

		conn.Close()
	}()

	if err := conn.Handshake(); err != nil {
		log.Error("handshake error %s", err.Error())
		c.Close()
		return
	}

	conn.Run()
}
