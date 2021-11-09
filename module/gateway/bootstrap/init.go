package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	conf2 "github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/logging"
)

func Init(path string) {
	InitApplication()
	conf2.Init(path)
	logging.InitLogging(conf2.LogConfig.Path, conf2.LogConfig.Level)
	if !conf2.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	model.Init()
}
