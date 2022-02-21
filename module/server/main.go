package main

import (
	"flag"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/module/server/g"
	"github.com/wenchangshou/vd-node-manage/module/server/http"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
	"github.com/wenchangshou/vd-node-manage/module/server/rpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		err error
	)
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	if *version {
		fmt.Printf("server %s version %s,build %s\n", "", "", "")
		os.Exit(0)
	}
	g.ParseConfig(*cfg)
	if err = model.InitDatabase(); err != nil {
		log.Fatalf("init database error:%s", err.Error())
	}
	if err := g.InitEvent(g.Config().Event.Provider, g.Config().Event.Arguments); err != nil {
		log.Fatalf("new event fail:" + err.Error())
	}
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
