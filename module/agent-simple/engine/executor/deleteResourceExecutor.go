package executor

import (
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/model"
	IService "github.com/wenchangshou/vd-node-manage/module/agent-simple/service"
	"path"

	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/zutil"
)

type DeleteResourceExecutor struct {
	Resource        *model.ResourceInfo
	DeviceService   IService.DeviceService
	eventService    IService.EventService
	ResourceService IService.ResourceService
}

func (executor *DeleteResourceExecutor) Execute() error {
	cfg := g.Config()
	resourcePath := path.Join(cfg.Resource.Directory, "resource", fmt.Sprintf("%d-%s", executor.Resource.ID, executor.Resource.Name))
	zutil.IsExistDelete(resourcePath)
	executor.DeviceService.DeleteComputerResource(executor.Resource.ID)
	return nil
	//return executor.ComputerService.DeleteComputerResource(executor.Option.ID)
}
func (executor *DeleteResourceExecutor) Cancel() error {
	return nil
}
func (executor *DeleteResourceExecutor) Verification(_ string) bool {
	return true
}
func (executor *DeleteResourceExecutor) SubscribeNotifyStatusChange(func(string, int, string)) {

}

// BindOption 检验任务参数
func (executor *DeleteResourceExecutor) BindOption(_ interface{}) error {
	return nil
}
