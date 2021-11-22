package bootstrap

import (
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent/g"
	"path"

	"github.com/wenchangshou2/zutil"
)

func Init() {
	InitApplication()
	cfg := g.Config()
	logging.InitLogging(cfg.Log.Path, cfg.Log.Level)
	zutil.IsNotExistMkDir(cfg.Resource.Directory)
	zutil.IsNotExistMkDir(cfg.Resource.Tmp)
	applicationPath := path.Join(cfg.Resource.Directory, "application")
	resourcePath := path.Join(cfg.Resource.Directory, "resource")
	zutil.IsNotExistMkDir(applicationPath)
	zutil.IsNotExistMkDir(resourcePath)
}
