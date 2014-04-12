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

var configJson = []byte(`
{
    "addr" : "127.0.0.1:4000",
    "user": "qing",
    "password": "admin",

    "nodes" : [
        {
            "name" : "node1",
            "mode" : "master",
            "switch_after_noalive": 300,
            "backends" : [
                {
                    "addr" : "127.0.0.1:3306",
                    "user" : "qing",
                    "password": "admin",
                    "db" : "mixer",
                    "idle_conns" : 32
                }
            ]
        },
        {
            "name" : "node2",
            "mode" : "slave",
            "switch_after_noalive": 60,

            "backends" : [
                {
                    "addr" : "127.0.0.1:3306",
                    "user" : "qing",
                    "password": "admin",
                    "db" : "mixer",
                    "idle_conns" : 32
                }
            ]
        }
    ],

    "schemas" : [
        {
            "db" : "mixer",
            "nodes" : ["node1", "node2"]
        }
    ]
}
	`)

func newTestServer() *Server {
	f := func() {
		cfg, err := newConfigJson(configJson)
		if err != nil {
			println(err.Error())
			panic(err)
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
		var cfg = []byte(`
			{
			    "addr" : "127.0.0.1:4000",
                "user" : "qing",
                "password": "admin",
                "db" : "mixer",
                "idle_conns" : 4
			}
			`)
		var err error
		testDB, err = NewDB(cfg)

		if err != nil {
			println(err.Error())
			panic(err)
		}
	}

	testDBOnce.Do(f)
	return testDB
}

func newTestDBConn() *SqlConn {
	db := newTestDB()

	c, err := NewSqlConn(db)

	if err != nil {
		println(err.Error())
		panic(err)
	}

	return c
}

func TestServer(t *testing.T) {
	newTestServer()
}
