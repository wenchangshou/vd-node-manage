package main

import (
	"flag"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/common/logging"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/buff"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/cron"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/engine"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/http"
	"github.com/wenchangshou2/zutil"
	"log"
	"path"
	"time"
)

var (
	confPath      string
	installFlag   bool
	uninstallFlag bool
)

func first() chan bool {
	r := make(chan bool)
	go func() {
		timeTicker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-timeTicker.C:
				if g.Config().Server.Register {
					r <- true
				}
			}
		}
	}()
	return r
}

// mkdirWorkDir 创建工作目录
func mkdirWorkDir(conf *g.GlobalConfig) {
	logging.InitLogging(conf.Log.Path, conf.Log.Level)
	// 创建工作目录
	ap := path.Join(conf.Resource.Directory, "application")
	rp := path.Join(conf.Resource.Directory, "resource")
	zutil.IsNotExistMkDir(conf.Resource.Directory)
	zutil.IsNotExistMkDir(conf.Resource.Tmp)
	zutil.IsNotExistMkDir(ap)
	zutil.IsNotExistMkDir(rp)
}
func InitSystemInfo(cfg *string, hardware *string) {
	var (
		err error
	)
	buff.InitGlobalBuffer()
	g.InitApplication()
	g.ParseConfig(*cfg)
	conf := g.Config()
	mkdirWorkDir(conf)
	g.ParseHardware(*hardware)
	go http.Start()
	<-first()
	g.InitLocalIp()
	g.InitRpcClients()
	cDriver, err := cache.InitCache("redis", g.Config().Server.RedisAddress, "", 0)
	if err != nil {
		log.Fatalln("初始化cache模块失败:" + err.Error())
	}
	cron.ReportDeviceStatus()
	if err = engine.InitSchedule(conf, cDriver); err != nil {
		log.Fatalln("初始化调度失败")
	}
}

type service struct {
}

func main() {

	cfg := flag.String("c", zutil.RelativePath("cfg.json"), "configuration file")
	hardware := flag.String("d", zutil.RelativePath("hardware.data"), "hardware file")
	flag.BoolVar(&installFlag, "i", false, "install program")
	flag.BoolVar(&uninstallFlag, "u", false, "uninstall program")
	flag.Parse()
	InitSystemInfo(cfg, hardware)
	select {}
	// Init
}
