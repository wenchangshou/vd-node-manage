package main

import (
	"flag"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/cron"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/http"
	"github.com/wenchangshou2/zutil"
)

var (
	confPath      string
	installFlag   bool
	uninstallFlag bool
)

// waitRegister 等待注册
func waitRegister() <-chan bool {
	result := make(chan bool)
	go func() {
		timeTicker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-timeTicker.C:
				if g.Config().Server.Register {
					result <- true
				}
			}
		}
	}()

	return result
}

func main() {
	var (
		err error
	)
	cfg := flag.String("c", zutil.RelativePath("cfg.json"), "configuration file")
	hardware := flag.String("d", zutil.RelativePath("hardware.data"), "hardware file")
	flag.BoolVar(&installFlag, "i", false, "install program")
	flag.BoolVar(&uninstallFlag, "u", false, "uninstall program")
	flag.Parse()
	// Init
	g.InitApplication()
	g.ParseConfig(*cfg)
	conf := g.Config()
	logging.InitLogging(conf.Log.Path, conf.Log.Level)
	// 创建工作目录
	ap := path.Join(conf.Resource.Directory, "application")
	rp := path.Join(conf.Resource.Directory, "resource")
	zutil.IsNotExistMkDir(conf.Resource.Directory)
	zutil.IsNotExistMkDir(conf.Resource.Tmp)
	zutil.IsNotExistMkDir(ap)
	zutil.IsNotExistMkDir(rp)
	g.ParseHardware(*hardware)
	go http.Start()
	result := waitRegister()
	<-result
	fmt.Println("注册成功")
	g.InitLocalIp()
	g.InitRpcClients()
	cron.ReportDeviceStatus()
	if err = engine.InitSchedule(conf); err != nil {
		log.Fatalln("初始化调度失败")
	}
	select {}
}
