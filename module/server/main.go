package main

import (
	"flag"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/server/g"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	if *version {
		fmt.Printf("server %s version %s,build %s\n", "", "", "")
		os.Exit(0)
	}
	g.ParseConfig(*cfg)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()
	select {}
}
