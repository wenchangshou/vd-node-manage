package main

import (
	"flag"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/module/server/event"
	"github.com/wenchangshou/vd-node-manage/module/server/g"
	"github.com/wenchangshou/vd-node-manage/module/server/http"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
	"github.com/wenchangshou/vd-node-manage/module/server/rpc"
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
	model.InitDatabase()
	event.InitEvent(g.Config().Cache)
	cache.InitCache("redis", g.Config().Cache.Addr, g.Config().Cache.Passwd, g.Config().Cache.DB)
	go http.Start()
	go rpc.Start()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	select {}
}
