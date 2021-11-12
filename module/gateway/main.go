package main

import (
	"flag"
	discover "github.com/wenchangshou2/vd-node-manage/common/discovery"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/bootstrap"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/g"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/routers"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server"
	"github.com/wenchangshou2/vd-node-manage/zebus"
	"net"
)

var (
	confPath string
)

func init() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	hardware:=flag.String("d","hardware.data","hardware file")
	flag.Parse()
	g.ParseConfig(*cfg)
	g.ParseHardware(*hardware)
	bootstrap.Init()
}
func main() {

	invention, _ := discover.NewInvention(net.IPv4zero, 8889)
	go invention.Server(e.ServerInfo{Ip: g.Config().Server.IP, Port: g.Config().Server.Port})
	api := routers.InitRouter()
	go rpcServer.InitRpc(":10051")
	zebus.InitZebus(g.Config().Zebus.IP, g.Config().Zebus.HttpPort, g.Config().Zebus.WsPort)
	if err := api.Run(g.Config().System.Listen); err != nil {
		logging.GLogger.Warn("无法监听[" + g.Config().System.Listen + "]" + "," + err.Error())
		panic(err)
	}
}
