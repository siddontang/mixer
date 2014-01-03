package proxy

import (
	"github/siddontang/golib/log"
)

type Proxy struct {
	cfg *Config
}

func NewProxy(cfg *Config) *Proxy {
	p := new(Proxy)

	p.cfg = cfg

	return p
}
