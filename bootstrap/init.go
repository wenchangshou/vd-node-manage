package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
)

func Init(path string) {
	InitApplication()
	conf.Init(path)
	logging.InitLogging(conf.LogConfig.Path, conf.LogConfig.Level)
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	model.Init()
}
