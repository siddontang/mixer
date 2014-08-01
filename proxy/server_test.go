package proxy

import (
	"github.com/siddontang/mixer/client"
	"sync"
	"testing"
	"time"
)

var testServerOnce sync.Once
var testServer *Server
var testDBOnce sync.Once
var testDB *client.DB

var testConfigData = []byte(`
addr : 127.0.0.1:4000
user : root
password : 

nodes :
- 
    name : node1 
    down_after_noalive : 300
    idle_conns : 16
    rw_split: true
    user: root
    password:
    master : 127.0.0.1:3306
    master_backup : 
    slave : 

schemas :
-
    db : mixer 
    nodes: [node1]

rules:
-   db: mixer
    table: 
    key:
    nodes: node1
    type: default
`)

func newTestServer() *Server {
	f := func() {
		cfg, err := ParseConfigData(testConfigData)
		if err != nil {
			println(err.Error())
			panic(err)
		}

		testServer, err = NewServer(cfg)
		if err != nil {
			println(err.Error())
			panic(err)
		}

		go testServer.Run()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer
}

func newTestDB() *client.DB {
	newTestServer()

	f := func() {
		cfg := new(client.Config)
		cfg.Addr = "127.0.0.1:4000"
		cfg.User = "root"
		cfg.DB = "mixer"
		cfg.IdleConns = 4

		var err error
		testDB, err = client.Open(cfg)

		if err != nil {
			println(err.Error())
			panic(err)
		}
	}

	testDBOnce.Do(f)
	return testDB
}

func newTestDBConn() *client.SqlConn {
	db := newTestDB()

	c, err := db.GetConn()

	if err != nil {
		println(err.Error())
		panic(err)
	}

	return c
}

func TestServer(t *testing.T) {
	newTestServer()
}
