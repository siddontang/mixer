package main

import (
	"flag"
	"strings"
	"github.com/siddontang/mixer/config"
	"github.com/siddontang/mixer/proxy"
	"github.com/siddontang/go-log/log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var configFile = flag.String("config", "/etc/mixer.conf", "mixer proxy config file")
var logLevel *string = flag.String("l", "[debug|info|warn|error]", "log level")

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
	
	if *logLevel != "" {
		setLogLevel(*logLevel)
	} else {
		setLogLevel(cfg.LogLevel)
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

func setLogLevel(level string){
	switch strings.ToLower(level) {
	case "debug":log.SetLevel(log.LevelDebug)
	case "info":log.SetLevel(log.LevelInfo)
	case "warn":log.SetLevel(log.LevelWarn)
	case "error":log.SetLevel(log.LevelError)
	default:
		log.SetLevel(log.LevelError)
	}
}
