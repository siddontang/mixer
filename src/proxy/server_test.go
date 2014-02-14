package proxy

import (
	. "mysql"
	"sync"
	"testing"
	"time"
)

var testServerOnce sync.Once
var testServer *Server
var testDBOnce sync.Once
var testDB *DB

func newTestServer() *Server {
	f := func() {
		cfg := new(config)

		cfg.Addr = "127.0.0.1:4000"
		cfg.User = "root"
		cfg.Password = ""

		cfg.Nodes = []configDataNode{
			configDataNode{"node1", "127.0.0.1:3306", "root", "", "mixer", "master", 4},
			configDataNode{"node2", "127.0.0.1:3306", "root", "", "mixer", "slave", 4},
		}

		cfg.Schemas = []configSchema{
			configSchema{"mixer", []string{"node1", "node2"}, true},
		}

		testServer = newServer(cfg)

		go testServer.Start()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer
}

func newTestDB() *DB {
	newTestServer()

	f := func() {
		testDB = NewDB("127.0.0.1:4000", "root", "", "mixer", 16)
	}

	testDBOnce.Do(f)
	return testDB
}

func TestServer(t *testing.T) {
	newTestServer()
}
