package proxy

import (
	"github.com/siddontang/golib/log"
	"net"
)

type Server struct {
	cfg *Config

	addr     string
	user     string
	password string

	running bool
}

func NewServer(cfg *Config) *Server {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.ConfigServer.Addr
	s.user = cfg.ConfigServer.User
	s.password = cfg.ConfigServer.Password

	return s
}

func (s *Server) Start() error {
	log.Info("start listen %s", s.addr)

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Error("listen error %s", err.Error())
		return err
	}

	s.running = true

	for s.running {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("accept error %s", err.Error())
			continue
		}

		go s.onConn(conn)
	}

	log.Info("stop listen")
}

func (s *Server) Stop() {
	s.running = false
}

func (s *Server) onConn(c net.Conn) {
	conn := NewClientConn(s, c)

	if err := conn.Handshake(); err != nil {
		log.Error("handshake error %s", err.Error())
		return
	}
}
