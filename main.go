package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wbwax/logger"
	"github.com/wbwax/wax-engine/g"
	"github.com/wbwax/wax-engine/services"
)

var (
	appName    string
	appVersion string
	gitCommit  string
)

func main() {
	g.UpdateLDFlags(appName, appVersion, gitCommit)
	cfgFile := flag.String("c", "conf/cfg.example.yaml", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	if *version {
		fmt.Printf("app name: %s, app version: %s, git commit: %s\n", g.AppName, g.AppVersion, g.GitCommit)
		os.Exit(0)
	}
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// load config
	if err := g.LoadServerConfig(*cfgFile); err != nil {
		fmt.Println("failed to load config, err:", err.Error())
		os.Exit(1)
	}

	// init logger
	logCfg := logger.Config{
		MaxSize:    g.Config.Log.MaxSize,
		MaxAge:     g.Config.Log.MaxAge,
		MaxBackups: g.Config.Log.MaxBackups,
		Level:      g.Config.Log.Level,
		Path:       g.Config.Log.Path,
		Encoding:   g.Config.Log.Encoding,
	}
	logger.Init(logCfg)
	logger.Infof("msg=%s||level=%s", "succeed to init logger", g.Config.Log.Level)
	defer logger.Sync() // flush buffer, if any

	go services.StartHTTP(g.Config.HTTP)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig // blocked
	fmt.Println("exited by signal:", s)
	logger.Infof("msg=%s||signal=%+v", "exited by signal", s)
	fmt.Println(1)
	os.Exit(0)
}
