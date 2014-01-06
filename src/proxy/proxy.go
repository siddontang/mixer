package proxy

import (
	"github.com/siddontang/golib/log"
)

type Proxy struct {
	cfg *Config
}

func NewProxy(cfg *Config) *Proxy {
	log.Info("NewProxy")
	p := new(Proxy)

	p.cfg = cfg

	return p
}
