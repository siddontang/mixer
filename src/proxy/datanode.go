package proxy

import (
	"container/list"
	"github.com/siddontang/golib/log"
	"strings"
	"sync"
)

const (
	MASTER_MODE byte = 0
	SLAVE_MODE  byte = 1
)

type DataNode struct {
	cfg *Config

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

func NewDataNode(cfg *Config, cfgNode *ConfigDataNode) *DataNode {
	dn := new(DataNode)

	dn.cfg = cfg
	dn.name = cfgNode.Name
	dn.addr = cfgNode.Addr
	dn.user = cfgNode.User
	dn.password = cfgNode.Password
	dn.db = cfgNode.DB

	dn.maxIdleConns = cfg.MaxIdleConns

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

func (dn *DataNode) PopConn() (*ProxyConn, error) {
	var c *ProxyConn

	dn.lock.Lock()
	if v := dn.conns.Back(); v != nil {
		dn.conns.Remove(v)
		c = v.Value.(*ProxyConn)
	}
	dn.lock.Unlock()

	if c != nil {
		if err := c.Ping(); err == nil {
			//connection has alive
			return c, nil
		}
	}

	c = NewProxyConn()

	if err := c.Connect(dn.addr, dn.user, dn.password, dn.db); err != nil {
		log.Error("connect %s node error %s", dn.name, err.Error())
		return nil, err
	}

	return c, nil
}

func (dn *DataNode) PushConn(c *ProxyConn) {
	dn.lock.Lock()

	if dn.conns.Len() > dn.maxIdleConns {
		dn.conns.Remove(dn.conns.Front())
	}

	dn.conns.PushBack(c)

	dn.lock.Unlock()
}

func (dn *DataNode) run() {
	dn.alive = true

	//to do
	//1 check connection alive
	//2 check remove mysql server alive
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

func NewDataNodes(cfg *Config) DataNodes {
	dns := make(DataNodes, len(cfg.DataNodes))

	for _, v := range cfg.DataNodes {
		dns[v.Name] = NewDataNode(cfg, &v)
	}

	return dns
}
