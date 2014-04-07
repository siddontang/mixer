package main

import (
	"flag"
	"github.com/siddontang/mixer/go/proxy"
)

var config = flag.String("config", "/etc/mixer/proxy.json", "config directory")

func main() {
	flag.Parse()

	s := proxy.NewServer(*config)
	s.Start()
}
