package main

import (
	"flag"
	"github.com/siddontang/mixer/config"
	"github.com/siddontang/mixer/proxy"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var configFile = flag.String("config", "/etc/mixer.conf", "mixer proxy config file")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if len(*configFile) == 0 {
		println("must use a config file")
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		println(err.Error())
		return
	}

	var svr *proxy.Server
	svr, err = proxy.NewServer(cfg)
	if err != nil {
		println(err.Error())
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-sc

		svr.Close()
	}()

	svr.Run()
}
