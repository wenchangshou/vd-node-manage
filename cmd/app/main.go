package main

import (
	"flag"

	"github.com/wenchangshou2/vd-node-manage/bootstrap"
	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"github.com/wenchangshou2/vd-node-manage/routers"
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
	api := routers.InitRouter()
	if err := api.Run(conf.SystemConfig.Listen); err != nil {
		logging.G_Logger.Warn("无法监听[" + conf.SystemConfig.Listen + "]" + "," + err.Error())
	}
}
