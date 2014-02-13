package proxy

import (
	"lib/log"
	"mysql"
	"strings"
	"time"
)

const (
	MASTER_MODE byte = 0
	SLAVE_MODE  byte = 1
)

type DataNode struct {
	server *Server
	cfg    *Config

	name string

	db *mysql.DB

	mode byte

	alive bool
}

func NewDataNode(server *Server, cfgNode *ConfigDataNode) *DataNode {
	dn := new(DataNode)

	dn.server = server
	dn.cfg = server.cfg

	dn.db = mysql.NewDB(cfgNode.Addr, cfgNode.User, cfgNode.Password, cfgNode.DB, server.cfg.MaxIdleConns)

	switch strings.ToLower(cfgNode.Mode) {
	case "master":
		dn.mode = MASTER_MODE
	case "slave":
		dn.mode = SLAVE_MODE
	default:
		log.Error("invalid node mode %s, use master instead", cfgNode.Mode)
		dn.mode = MASTER_MODE
	}

	go dn.run()

	return dn
}

func (dn *DataNode) DB() *mysql.DB {
	return dn.db
}

func (dn *DataNode) run() {
	dn.alive = true

	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	var errNum int = 0

	t := time.NewTicker(3 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if err := dn.db.Ping(); err != nil {
				log.Error("ping error %s", err.Error())
				errNum++
			} else {
				errNum = 0
				dn.alive = true
			}

			if errNum > 3 {
				log.Error("check alive 3 failed, disable alive")
				dn.alive = false
			}
		}
	}
}

func (dn *DataNode) IsAlive() bool {
	return dn.alive
}

type DataNodes map[string]*DataNode

func (dns DataNodes) GetNode(name string) *DataNode {
	if dn, ok := dns[name]; ok {
		return dn
	} else {
		return nil
	}
}

func NewDataNodes(server *Server) DataNodes {
	cfg := server.cfg

	dns := make(DataNodes, len(cfg.DataNodes))

	for _, v := range cfg.DataNodes {
		dns[v.Name] = NewDataNode(server, &v)
	}

	return dns
}
