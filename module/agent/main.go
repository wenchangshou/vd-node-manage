package main

import (
	"flag"
	"fmt"
	SystemService "github.com/kardianos/service"
	"github.com/wenchangshou2/vd-node-manage/module/agent/bootstrap"
	"github.com/wenchangshou2/vd-node-manage/module/agent/engine"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/invention"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/logging"
	"github.com/wenchangshou2/zutil"
)

var (
	confPath      string
	installFlag   bool
	uninstallFlag bool
)

type program struct{}

func (p *program) Start(s SystemService.Service) error {
	logging.G_Logger.Info("启动程序")
	go p.run()
	return nil
}
func (p *program) run() {
	var (
		serverInfo *e.ServerInfo
	)
	// 是否启用自动发现
	if conf.SystemConfig.Mode == "auto" {
		logging.G_Logger.Info("等待服务端响应注册请求")
		invention, err := invention.NewInvention("", 8889)
		if err != nil {
			logging.G_Logger.Error(fmt.Sprintf("初始化发现服务失败：%v\n", err))
			panic(err)
		}
		serverInfo, err = invention.WaitRegister()
		if err != nil {
			logging.G_Logger.Error(fmt.Sprintf("等待注册失败:%s\n", err.Error()))
			panic(err)
		}

	} else {
		// 读取配置文件
		serverInfo.Ip = conf.SystemConfig.Ip
		serverInfo.Port = int(conf.SystemConfig.Port)
	}

	fmt.Printf("接收到服务端信息,ip:%s,port:%d\n", serverInfo.Ip, serverInfo.Port)
	// service.InitService(serverInfo.Ip, serverInfo.Port)
	// engine.InitTaskExecute(1, fmt.Sprintf("%s:%d", serverInfo.Ip, serverInfo.Port), conf.RpcConfig.Address)
	err := engine.InitSchedule(serverInfo, conf.SystemConfig.HashIDSalt)
	if err != nil {
		return
	}
	// engine.InitRpcPubSub(conf.RpcConfig.Address, serverInfo.Ip)
}
func (p *program) Stop(s SystemService.Service) error {
	return nil
}

func init() {
	flag.StringVar(&confPath, "c", zutil.RelativePath("conf.ini"), "配置文件路径")
	flag.BoolVar(&installFlag, "i", false, "install program")
	flag.BoolVar(&uninstallFlag, "u", false, "uninstall program")
	flag.Parse()
	bootstrap.Init(confPath)
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
		logging.G_Logger.Error("创建服务失败:" + err.Error())
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
