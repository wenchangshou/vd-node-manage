package main

import (
	"flag"
	"fmt"
	SystemService "github.com/kardianos/service"
	discover "github.com/wenchangshou2/vd-node-manage/common/discovery"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/cron"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/http"
	"github.com/wenchangshou2/zutil"
	"net"
	"path"
)

var (
	confPath      string
	installFlag   bool
	uninstallFlag bool
)

type program struct{}

func (p *program) Start(SystemService.Service) error {
	fmt.Println("启动程序")
	go p.run()
	return nil
}
func (p *program) run() {
	hardware := g.Hardware()
	g.InitLocalIp()
	g.InitRpcClients()
	go http.Start()
	cfg := g.Config()
	cron.ReportDeviceStatus()
	// 是否启用自动发现
	if cfg.Server.Mode == "auto" {
		invention, err := discover.NewInvention(net.IPv4zero, 0)
		if err != nil {
			logging.GLogger.Error(fmt.Sprintf("初始化发现服务失败：%v\n", err))
			return
		}
		serverInfo, err := invention.WaitRegister()
		if err != nil {
			logging.GLogger.Error(fmt.Sprintf("等待注册失败:%s\n", err.Error()))
			return
		}
		cfg.Server.Address = fmt.Sprintf("%s:%d", serverInfo.Ip, serverInfo.Port)
	}
	err := engine.InitSchedule(g.Config().Server.HttpAddress,g.Config().Server.RpcAddress, hardware.ID)
	if err != nil {
		return
	}
}
func (p *program) Stop(SystemService.Service) error {
	return nil
}

func init() {
	cfg := flag.String("c", zutil.RelativePath("cfg.json"), "configuration file")
	hardware := flag.String("d", zutil.RelativePath("hardware.data"), "hardware file")
	flag.BoolVar(&installFlag, "i", false, "install program")
	flag.BoolVar(&uninstallFlag, "u", false, "uninstall program")
	flag.Parse()

	// Init
	g.InitApplication()
	g.ParseConfig(*cfg)
	conf:=g.Config()
	logging.InitLogging(conf.Log.Path,conf.Log.Level)
	// 创建工作目录
	ap:=path.Join(conf.Resource.Directory,"application")
	rp:=path.Join(conf.Resource.Directory,"resource")
	zutil.IsNotExistMkDir(conf.Resource.Directory)
	zutil.IsNotExistMkDir(conf.Resource.Tmp)
	zutil.IsNotExistMkDir(ap)
	zutil.IsNotExistMkDir(rp)
	g.ParseHardware(*hardware)
}
func main() {
	var (
		err error
	)
	svcConfig := &SystemService.Config{
		Name:        "vd client",
		DisplayName: "vd客户端",
		Description: "",
	}
	prg := &program{}
	s, err := SystemService.New(prg, svcConfig)
	if err != nil {
		logging.GLogger.Error("创建服务失败:" + err.Error())
	}
	if installFlag {
		err := s.Install()
		if err != nil {
			fmt.Println("服务安装失败:" + err.Error())
			return
		}
		fmt.Println("服务安装成功")
		return
	}
	if uninstallFlag {
		err := s.Install()
		if err != nil {
			fmt.Println("服务卸载失败:" + err.Error())
			return
		}
		fmt.Println("服务卸载成功")
		return
	}
	err = s.Run()
	if err != nil {
		fmt.Println("err", err)
	}

}
