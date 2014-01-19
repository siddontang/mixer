package proxy

import (
	"github.com/siddontang/golib/log"
	"strings"
)

const (
	MASTER_MODE byte = 0
	SLAVE_MODE  byte = 1
)

type DataNode struct {
	cfg      *Config
	name     string
	addr     string
	user     string
	password string
	mode     byte
}

func NewDataNode(cfg *Config, name, addr, user, password, mode string) *DataNode {
	dn := new(DataNode)

	dn.cfg = cfg
	dn.name = name
	dn.addr = addr
	dn.user = user
	dn.password = password

	switch strings.ToLower(mode) {
	case "master":
		dn.mode = MASTER_MODE
	case "slave":
		dn.mode = SLAVE_MODE
	default:
		log.Error("invalid node mode %s, use master instead", mode)
		dn.mode = MASTER_MODE
	}

	return dn
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
	dns := make(DataNodes, len(cfg.ConfigSchema.DataNodes))

	for _, v := range cfg.ConfigSchema.DataNodes {
		dns[v.Name] = NewDataNode(cfg, v.Name, v.Addr, v.User, v.Password, v.Mode)
	}

	return dns
}
