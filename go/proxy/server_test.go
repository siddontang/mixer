package proxy

import (
	. "github.com/siddontang/mixer/go/mysql"
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
		cfg.User = "qing"
		cfg.Password = "admin"

		cfg.Nodes = []configDataNode{
			configDataNode{"node1", []string{"qing:admin@127.0.0.1:3306/mixer"}, "master", 300, 4},
			configDataNode{"node2", []string{"qing:admin@127.0.0.1:3306/mixer"}, "slave", 60, 4},
		}

		cfg.Schemas = []configSchema{
			configSchema{"mixer", []string{"node1", "node2"}},
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
		testDB, _ = NewDB("qing:admin@127.0.0.1:4000/mixer", 16)
	}

	testDBOnce.Do(f)
	return testDB
}

func TestServer(t *testing.T) {
	newTestServer()
}
