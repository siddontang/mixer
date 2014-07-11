package main

import (
	"flag"
	"github.com/siddontang/mixer/server"
)

var config = flag.String("config", "/etc/mixer.json", "config directory")

func main() {
	flag.Parse()

	s := server.NewServer(*config)
	s.Start()
}
