package proxy

import (
	"fmt"
	"log"
	"net"
	"os"
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
		errLog("listen error %s", err.Error())
		return err
	}

	s.running = true

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			errLog("accept error %s", err.Error())
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
			errLog("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}

		conn.Close()
	}()

	if err := conn.Handshake(); err != nil {
		errLog("handshake error %s", err.Error())
		c.Close()
		return
	}

	conn.Run()
}

var (
	errLogger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
)

func errLog(format string, args ...interface{}) {
	f := fmt.Sprintf("[Error] [mixer.proxy] %s", format)
	s := fmt.Sprintf(f, args...)
	errLogger.Output(2, s)
}
