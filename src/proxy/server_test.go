package proxy

import (
	"sync"
	"testing"
	"time"
)

var testServerOnce sync.Once
var testServer *Server

func newTestServer() *Server {
	f := func() {
		cfg := new(Config)

		cfg.Addr = "127.0.0.1:4000"
		cfg.User = "qing"
		cfg.Password = "admin"
		cfg.MaxIdleConns = 4

		cfg.DataNodes = []ConfigDataNode{
			ConfigDataNode{"node1", "10.20.135.213:3306", "qing", "admin", "mixer", "master"},
			ConfigDataNode{"node2", "10.20.135.213:3306", "qing", "admin", "mixer", "slave"},
		}

		cfg.Schemas = []ConfigSchema{
			ConfigSchema{"mixer", []string{"node1", "node2"}, true},
		}

		testServer = NewServer(cfg)

		go testServer.Start()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer
}

func TestServer(t *testing.T) {
	newTestServer()
}
