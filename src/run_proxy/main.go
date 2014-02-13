package main

import (
	"flag"
	"proxy"
)

var configDir = flag.String("configDir", "../etc/proxy", "config directory")

func main() {
	flag.Parse()

	s := proxy.NewServer(*configDir)
	s.Start()
}
