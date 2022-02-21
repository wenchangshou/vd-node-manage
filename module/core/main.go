package main

import (
	"flag"
	"github.com/kardianos/service"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/common/logging"
	"github.com/wenchangshou/vd-node-manage/module/core/buff"
	"github.com/wenchangshou/vd-node-manage/module/core/cron"
	"github.com/wenchangshou/vd-node-manage/module/core/engine"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
	"github.com/wenchangshou/vd-node-manage/module/core/g/db"
	"github.com/wenchangshou/vd-node-manage/module/core/g/process"
	"github.com/wenchangshou/vd-node-manage/module/core/http"
	"github.com/wenchangshou2/zutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	confPath      string
	installFlag   bool
	uninstallFlag bool
	restartFlag   bool
	modeFlag      string
	cfg           string
	hardware      string
)
var logger service.Logger

func first() chan bool {
	r := make(chan bool)
	go func() {
		timeTicker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-timeTicker.C:
				if g.GetServerInfo().Register {
					r <- true
				}
			default:
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

type Program struct {
	schedule *engine.Schedule
	restart  chan bool
}

func (p *Program) start() {
	conf := g.Config()
	<-first()
	g.InitLocalIp()
	g.InitRpcClients()
	cDriver, err := cache.InitCache("redis", g.GetServerInfo().Redis.Address, "", 0)
	if err != nil {
		log.Fatalln("初始化cache模块失败:" + err.Error())
	}
	cron.ReportDeviceStatus()
	if p.schedule, err = engine.InitSchedule(conf, cDriver, p.restart); err != nil {
		log.Fatalln("初始化调度失败")
	}
}
func (p *Program) Start(_ service.Service) error {

	db.InitDB("data.db")
	process.GenerateProcess(modeFlag)
	buff.InitGlobalBuffer()
	g.InitApplication()
	g.ParseConfig(cfg)
	conf := g.Config()
	mkdirWorkDir(conf)
	g.ParseHardware(hardware)
	g.LoadServerInfoByDb()
	go http.Start(p.restart)
	go p.start()
	return nil
}
func (p *Program) Stop(_ service.Service) error {
	logger.Info("stop")
	if p.schedule != nil {
		p.schedule.Exit()
	}
	return nil
}
func GetCurrentPath() string {
	fullExecPath, _ := os.Executable()
	dir, _ := filepath.Split(fullExecPath)
	return dir
}
func main() {
	var (
		s   service.Service
		err error
	)
	os.Chdir(GetCurrentPath())
	svcConfig := &service.Config{
		Name:             "quickex-service",
		DisplayName:      "quickex-service",
		Description:      "快展后台服务",
		Arguments:        []string{"-m", "service"},
		WorkingDirectory: GetCurrentPath(),
	}
	restart := make(chan bool)
	svc := &Program{restart: restart}
	s, err = service.New(svc, svcConfig)
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	flag.StringVar(&cfg, "c", "cfg.json", "configuration file")
	flag.StringVar(&hardware, "d", "hardware.data", "hardware file")
	flag.BoolVar(&installFlag, "i", false, "install program")
	flag.BoolVar(&uninstallFlag, "u", false, "uninstall program")
	flag.BoolVar(&restartFlag, "s", false, "restart program")
	flag.StringVar(&modeFlag, "m", "console", "service model ")
	flag.Parse()
	if installFlag {
		err := s.Install()
		if err != nil {
			log.Println("安裝服務失敗:" + err.Error())
			return
		}
		log.Println("安裝服務成功")
		return
	}
	if uninstallFlag {
		err := s.Uninstall()
		if err != nil {
			log.Println("卸載服務失敗:" + err.Error())
			return
		}
		log.Println("卸載服務成功")
		return
	}
	if restartFlag {
		s.Restart()
		return
	}

	err = s.Run()
	if err != nil {
		log.Fatalln("啟動服務失敗:" + err.Error())
	}
}
