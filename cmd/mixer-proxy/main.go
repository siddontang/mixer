package main

import (
	"flag"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/mixer/config"
	"github.com/siddontang/mixer/proxy"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

var configFile *string = flag.String("config", "/etc/mixer.conf", "mixer proxy config file")
var logLevel *string = flag.String("log-level", "", "log level [debug|info|warn|error], default error")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if len(*configFile) == 0 {
		log.Error("must use a config file")
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		log.Error(err.Error())
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
		log.Error(err.Error())
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		log.Info("Got signal [%d] to exit.", sig)
		svr.Close()
	}()

	svr.Run()
}

func setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		log.SetLevel(log.LevelDebug)
	case "info":
		log.SetLevel(log.LevelInfo)
	case "warn":
		log.SetLevel(log.LevelWarn)
	case "error":
		log.SetLevel(log.LevelError)
	default:
		log.SetLevel(log.LevelError)
	}
}
