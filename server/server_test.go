package server

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

var configJson = []byte(`
{
    "addr" : "127.0.0.1:4000",
    "user": "root",
    "password": "",

    "nodes" : [
        {
            "name" : "node1",
            "mode" : "master",
            "switch_after_noalive": 300,
            "backends" : [
                {
                    "addr" : "127.0.0.1:3306",
                    "user" : "root",
                    "password": "",
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
                    "user" : "root",
                    "password": "",
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

func newTestDB() *client.DB {
	newTestServer()

	f := func() {
		var cfg = []byte(`
			{
			    "addr" : "127.0.0.1:4000",
                "user" : "root",
                "password": "",
                "db" : "mixer",
                "idle_conns" : 4
			}
			`)
		var err error
		testDB, err = client.NewDB(cfg)

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

	c, err := client.NewSqlConn(db)

	if err != nil {
		println(err.Error())
		panic(err)
	}

	return c
}

func TestServer(t *testing.T) {
	newTestServer()
}
