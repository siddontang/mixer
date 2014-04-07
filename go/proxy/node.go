package proxy

import (
	"container/list"
	"fmt"
	"github.com/siddontang/golib/log"
	. "github.com/siddontang/mixer/go/mysql"
	"strings"
	"sync"
	"time"
)

const (
	MASTER_MODE byte = 0
	SLAVE_MODE  byte = 1
)

type node struct {
	sync.Mutex

	server *Server
	cfg    *config

	Name string

	//current running db
	db *DB

	dbs *list.List

	switchAfterNoAlive time.Duration

	Mode byte
}

func newNode(server *Server, cfgNode *configDataNode) (*node, error) {
	n := new(node)

	n.Name = cfgNode.Name
	n.server = server
	n.cfg = server.cfg

	if len(cfgNode.DSN) == 0 {
		return nil, fmt.Errorf("no dsn set")
	}

	n.dbs = list.New()

	var err error
	var db *DB
	for _, dsn := range cfgNode.DSN {
		db, err = NewDB(dsn, cfgNode.MaxIdleConns)
		if err != nil {
			return nil, err
		}
		n.dbs.PushBack(db)
	}

	n.db = n.dbs.Front().Value.(*DB)

	if err != nil {
		return nil, err
	}

	n.switchAfterNoAlive = time.Duration(cfgNode.SwitchAfterNoAlive) * time.Second

	switch strings.ToLower(cfgNode.Mode) {
	case "master":
		n.Mode = MASTER_MODE
	case "slave":
		n.Mode = SLAVE_MODE
	default:
		log.Error("invalid node mode %s, use master instead", cfgNode.Mode)
		n.Mode = MASTER_MODE
	}

	go n.run()

	return n, nil
}

func (n *node) GetConn() (*Conn, error) {
	n.Lock()
	db := n.db
	n.Unlock()

	return db.GetConn()
}

func (n *node) run() {
	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	t := time.NewTicker(3 * time.Second)
	defer t.Stop()

	lastPing := time.Now().Unix()
	for {
		select {
		case <-t.C:
			n.Lock()
			db := n.db
			n.Unlock()

			if err := db.Ping(); err != nil {
				log.Error("ping %s error %s", db.Addr(), err.Error())
			} else {
				lastPing = time.Now().Unix()
				break
			}

			if time.Now().Unix()-lastPing > int64(n.switchAfterNoAlive) {
				log.Error("db %s not alive over %ds, switch another",
					db.Addr(), int64(n.switchAfterNoAlive/time.Second))

				n.switchOver()
			}
		}
	}
}

func (n *node) switchOver() {
	v := n.dbs.Front()
	n.dbs.Remove(v)

	db := v.Value.(*DB)

	db.Close()

	n.dbs.PushBack(db)

	db = n.dbs.Front().Value.(*DB)

	n.Lock()
	n.db = db
	n.Unlock()
}

type nodes map[string]*node

func (ns nodes) GetNode(name string) *node {
	if n, ok := ns[name]; ok {
		return n
	} else {
		return nil
	}
}

func newNodes(server *Server) nodes {
	cfg := server.cfg

	ns := make(nodes, len(cfg.Nodes))

	var err error
	var n *node
	for _, v := range cfg.Nodes {
		n, err = newNode(server, &v)
		if err != nil {
			log.Error("new node %s error %s", v.Name, err.Error())
		} else {
			ns[v.Name] = n
		}
	}

	return ns
}
