package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wbwax/wax-engine/g"
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
		fmt.Printf("app name: %s, app version: %s, git commit: %s\n", appName, appVersion, gitCommit)
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig // blocked
	fmt.Println("exited by signal:", s)
	os.Exit(0)
}
