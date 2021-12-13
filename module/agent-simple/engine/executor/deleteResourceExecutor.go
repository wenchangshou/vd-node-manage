package executor

import (
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/zutil"
	"path"
)

type DeleteResourceExecutor struct {
	Option DeleteOption
}

func (executor *DeleteResourceExecutor) Execute() error {
	cfg := g.Config()
	resourcePath := path.Join(cfg.Resource.Directory, "resource", executor.Option.File.GetResourcePath())
	zutil.IsExistDelete(resourcePath)
	return nil
	//return executor.ComputerService.DeleteComputerResource(executor.Option.ID)
}
func (executor *DeleteResourceExecutor) Cancel() error {
	return nil
}
func (executor *DeleteResourceExecutor) Verification(option string) bool {
	return true
}
func (executor *DeleteResourceExecutor) SubscribeNotifyStatusChange(func(string, int, string)) {

}

// BindOption 检验任务参数
func (executor *DeleteResourceExecutor) BindOption(option interface{}) error {
	return nil
}
