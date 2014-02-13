package proxy

import (
	"lib/log"
	. "mysql"
	"strings"
	"time"
)

const (
	MASTER_MODE byte = 0
	SLAVE_MODE  byte = 1
)

type node struct {
	server *Server
	cfg    *config

	Name string

	*DB

	Mode byte

	Alive bool
}

func newNode(server *Server, cfgNode *configDataNode) *node {
	n := new(node)

	n.Name = cfgNode.Name
	n.server = server
	n.cfg = server.cfg

	n.DB = NewDB(cfgNode.Addr, cfgNode.User, cfgNode.Password, cfgNode.DB, cfgNode.MaxIdleConns)

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

	return n
}

func (n *node) run() {
	n.Alive = true

	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	var errNum int = 0

	t := time.NewTicker(3 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if err := n.Ping(); err != nil {
				log.Error("ping error %s", err.Error())
				errNum++
			} else {
				errNum = 0
				n.Alive = true
			}

			if errNum > 3 {
				log.Error("check alive 3 failed, disable alive")
				n.Alive = false
			}
		}
	}
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

	for _, v := range cfg.Nodes {
		ns[v.Name] = newNode(server, &v)
	}

	return ns
}
