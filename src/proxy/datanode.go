package proxy

import (
	"container/list"
	"github.com/siddontang/golib/log"
	"mysql"
	"strings"
	"sync"
	"time"
)

const (
	MASTER_MODE byte = 0
	SLAVE_MODE  byte = 1
)

type DataNode struct {
	server *Server
	cfg    *Config

	name     string
	addr     string
	user     string
	password string
	db       string
	mode     byte

	maxIdleConns int

	lock  sync.Mutex
	conns *list.List

	alive bool
}

func NewDataNode(server *Server, cfgNode *ConfigDataNode) *DataNode {
	dn := new(DataNode)

	dn.server = server
	dn.cfg = server.cfg

	dn.name = cfgNode.Name
	dn.addr = cfgNode.Addr
	dn.user = cfgNode.User
	dn.password = cfgNode.Password
	dn.db = cfgNode.DB

	dn.maxIdleConns = server.cfg.MaxIdleConns

	dn.conns = list.New()

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

func (dn *DataNode) PopConn() (*mysql.Client, error) {
	var c *mysql.Client

	dn.lock.Lock()
	if v := dn.conns.Back(); v != nil {
		dn.conns.Remove(v)
		c = v.Value.(*mysql.Client)
	}
	dn.lock.Unlock()

	if c != nil {
		if err := c.Ping(); err == nil {
			//connection has alive
			return c, nil
		}
	}

	c = mysql.NewClient()

	if err := c.Connect(dn.addr, dn.user, dn.password, dn.db); err != nil {
		log.Error("connect %s node error %s", dn.name, err.Error())
		return nil, err
	}

	//we must always use autocommit
	if _, err := c.Exec("set autocommit = 1"); err != nil {
		log.Error("set autocommit error %s", err.Error())
		c.Close()

		return nil, err
	}

	return c, nil
}

func (dn *DataNode) PushConn(c *mysql.Client) {
	var closeConn *mysql.Client
	dn.lock.Lock()

	if dn.conns.Len() > dn.maxIdleConns {
		oldConn := dn.conns.Front()
		dn.conns.Remove(oldConn)

		closeConn = oldConn.Value.(*mysql.Client)
	}

	dn.conns.PushBack(c)

	dn.lock.Unlock()

	if closeConn != nil {
		closeConn.Close()
	}
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
			if c, err := dn.PopConn(); err != nil {
				log.Error("pop conn error %s", err.Error())
				errNum++
			} else {
				if err := c.Ping(); err != nil {
					log.Error("ping error %s", err.Error())
					errNum++
				} else {
					errNum = 0
					dn.alive = true
					dn.PushConn(c)
				}
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
