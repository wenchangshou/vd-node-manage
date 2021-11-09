package executor

import (
	"encoding/json"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/conf"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent/service"
	"path"

	"github.com/wenchangshou2/zutil"
)

type DeleteOption struct {
	ID   string  `json:"id"`
	File File `json:"file"`
}
type DeleteProjectExecutor struct {
	Option          DeleteOption
	ComputerService IService.ComputerService
	TaskID string
}

func (executor *DeleteProjectExecutor) Execute() error {
	projectPath := path.Join(conf.ResourceConfig.Directory, "application", executor.Option.File.Uuid)
	zutil.IsExistDelete(projectPath)
	return executor.ComputerService.DeleteComputerProject(executor.Option.ID)
}
func (executor *DeleteProjectExecutor) Cancel() error {
	return nil
}
func (executor *DeleteProjectExecutor) Verification(option string) bool {
	return true
}
func (executor *DeleteProjectExecutor) SubscribeNotifyStatusChange(func(string, int, string)) {

}

// Verification 检验任务参数
func (executor *DeleteProjectExecutor) BindOption(option string) error {
	err := json.Unmarshal([]byte(option), &executor.Option)
	return err
}
