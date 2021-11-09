package bootstrap

import (
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/logging"
	"path"

	"github.com/wenchangshou2/zutil"
)

func Init(confPath string) {
	InitApplication()
	err := conf.Init(confPath)
	if err != nil {
		panic(err)
	}
	logging.InitLogging(conf.LogConfig.Path, conf.LogConfig.Level)
	zutil.IsNotExistMkDir(conf.ResourceConfig.Directory)
	zutil.IsNotExistMkDir(conf.ResourceConfig.Tmp)
	applicationPath := path.Join(conf.ResourceConfig.Directory, "application")
	resourcePath := path.Join(conf.ResourceConfig.Directory, "resource")
	zutil.IsNotExistMkDir(applicationPath)
	zutil.IsNotExistMkDir(resourcePath)
}
