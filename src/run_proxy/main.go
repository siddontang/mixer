package main

import (
	"flag"
	"proxy"
)

var configDir = flag.String("configDir", "../etc", "config directory")

func main() {
	flag.Parse()
	cfg, err := proxy.NewConfig(*configDir)
	if err != nil {
		panic(err)
	}

	s := proxy.NewServer(cfg)
	s.Start()
}
