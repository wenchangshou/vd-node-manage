package main

import (
	"flag"
	"fmt"
	SystemService "github.com/kardianos/service"
	discover "github.com/wenchangshou2/vd-node-manage/common/discovery"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent/bootstrap"
	"github.com/wenchangshou2/vd-node-manage/module/agent/engine"
	"github.com/wenchangshou2/vd-node-manage/module/agent/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
	"github.com/wenchangshou2/zutil"
	"net"
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
	var (
		serverInfo *e.ServerInfo
	)
	hardware := g.Hardware()
	cfg := g.Config()
	// 是否启用自动发现
	if cfg.System.Mode == "auto" {
		invention, err := discover.NewInvention(net.IPv4zero, 0)
		if err != nil {
			logging.GLogger.Error(fmt.Sprintf("初始化发现服务失败：%v\n", err))
			panic(err)
		}
		serverInfo, err = invention.WaitRegister()
		if err != nil {
			logging.GLogger.Error(fmt.Sprintf("等待注册失败:%s\n", err.Error()))
			panic(err)
		}

	} else {
		// 读取配置文件
		serverInfo.Ip = cfg.System.IP
		serverInfo.Port = cfg.System.Port
	}

	fmt.Printf("接收到服务端信息,ip:%s,port:%d\n", serverInfo.Ip, serverInfo.Port)
	// service.InitService(serverInfo.Ip, serverInfo.Port)
	// engine.InitTaskExecute(1, fmt.Sprintf("%s:%d", serverInfo.Ip, serverInfo.Port), conf2.RpcConfig.Address)
	err := engine.InitSchedule(serverInfo, hardware.ID)
	if err != nil {
		return
	}
	// engine.InitRpcPubSub(conf2.RpcConfig.Address, serverInfo.Ip)
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
	g.ParseConfig(*cfg)
	g.ParseHardware(*hardware)
	bootstrap.Init()
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
