package server

import (
	"github.com/siddontang/go-log/log"
	"net"
	"runtime"
	"strings"
)

type Server struct {
	cfg *config

	addr     string
	user     string
	password string

	nodes   nodes
	schemas schemas

	running bool

	listener net.Listener
}

func newServer(cfg *config) *Server {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.Addr
	s.user = cfg.User
	s.password = cfg.Password

	s.nodes = newNodes(s)
	s.schemas = newSchemas(s, s.nodes)

	return s
}

func NewServer(cfgDir string) *Server {
	cfg, err := newConfig(cfgDir)
	if err != nil {
		panic(err)
		return nil
	}

	return newServer(cfg)
}

func (s *Server) Start() error {
	var err error

	netProto := "tcp"
	if strings.Contains(s.addr, "/") {
		netProto = "unix"
	}

	s.listener, err = net.Listen(netProto, s.addr)

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
	conn := newconn(s, c)

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
