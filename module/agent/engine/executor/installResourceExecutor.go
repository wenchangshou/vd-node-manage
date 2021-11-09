package executor

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/conf"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent/service"
	"github.com/wenchangshou2/vd-node-manage/module/agent/util"
	"path"

)

type InstallResourceExecutor struct {
	Option          InstallResourceOption
	HttpRequestUri string
	taskService IService.TaskService
	computerService IService.ComputerService
	Mac             string
	taskID string
}
type InstallResourceOption struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Uri  string
}

func (executor *InstallResourceExecutor) Execute() error {
	requestUrl := "http://" + executor.HttpRequestUri + "/" + executor.Option.Uri
	dstPath := path.Join(conf.ResourceConfig.Directory, "resource/")
	err := util.DownloadFile(requestUrl, dstPath, executor.Option.ID+"-"+executor.Option.Name, func(length, downLen int64) {
		fmt.Printf("download:len:%d,downLen:%d\n", length, downLen)
	})
	if err != nil {
		executor.taskService.SetTaskItemStatus([]string{executor.taskID},ERROR)
		//executor.NotifyEvent(executor.TaskID, ERROR, "下载文件失败")
		return err
	}
	err = executor.computerService.AddComputerResource(executor.Option.ID)
	if err != nil {
		executor.taskService.SetTaskItemStatus([]string{executor.taskID},ERROR)
		return err
	}
	executor.taskService.SetTaskItemStatus([]string{executor.taskID},DONE)
	return nil
}

func (executor *InstallResourceExecutor) Cancel() error {
	return nil
}
func (executor *InstallResourceExecutor) Verification(option string) bool {
	err := json.Unmarshal([]byte(option), &executor.Option)
	return err == nil
}

// BindOption  检验任务参数
func (executor *InstallResourceExecutor) BindOption(option string) error {
	err := json.Unmarshal([]byte(option), &executor.Option)
	return err
}
