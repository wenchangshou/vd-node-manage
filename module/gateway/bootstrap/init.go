package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/g"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

func Init() {
	InitApplication()
	logging.InitLogging(g.Config().Log.Path, g.Config().Log.Level)
	if !g.Config().System.Debug{
		gin.SetMode(gin.ReleaseMode)
	}
	model.Init()
}
