package main

import (
	"flag"
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/bootstrap"
	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/pkg/discovery"
	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"github.com/wenchangshou2/vd-node-manage/routers"
	rpcServer "github.com/wenchangshou2/vd-node-manage/rpc/server"
	"github.com/wenchangshou2/vd-node-manage/zebus"
	"github.com/wenchangshou2/zutil"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", zutil.RelativePath("conf.ini"), "配置文件路径")
	flag.Parse()
	bootstrap.Init(confPath)
}
func main() {
	if err := discovery.InitDiscovery("0.0.0.0", 8889); err != nil {
		fmt.Println("初始化discovery失败")
		return
	}
	api := routers.InitRouter()
	go rpcServer.InitRpc(":10051")
	zebus.InitZebus(conf.ZebusConfig.Ip, conf.ZebusConfig.HttpPort, conf.ZebusConfig.WsPort)
	if err := api.Run(conf.SystemConfig.Listen); err != nil {
		logging.G_Logger.Warn("无法监听[" + conf.SystemConfig.Listen + "]" + "," + err.Error())
	}
}
