package executor

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"

	"github.com/mitchellh/mapstructure"
	"github.com/wenchangshou2/vd-node-manage/common/file"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent-simple/service"
)

type InstallResourceExecutor struct {
	Option         InstallResourceOption
	HttpRequestUri string
	eventService   IService.EventService
	DeviceService  IService.DeviceService
	Mac            string
	taskID         uint
}
type InstallResourceOption struct {
	ResourceID uint   `json:"resource_id"`
	Name       string `json:"name"`
	Uri        string `json:"uri"`
}

func (executor *InstallResourceExecutor) Execute() error {
	cfg := g.Config()
	//requestUrl := "http://" + executor.HttpRequestUri + "/" + executor.Option.Uri
	dstPath := path.Join(cfg.Resource.Directory, "resource/")
	err := file.DownloadFile(executor.Option.Uri, dstPath, fmt.Sprintf("%d-%s", executor.Option.ResourceID, executor.Option.Name), func(length, downLen int64) {
		fmt.Printf("download:len:%d,downLen:%d\n", length, downLen)
	})
	if err != nil {
		return fmt.Errorf("%s:%v", "下载文件失败", err)
	}
	if err = executor.DeviceService.AddComputerResource(executor.Option.ResourceID); err != nil {
		return errors.New("添加设备资源失败:" + err.Error())
	}
	fmt.Println("err", err)
	//if err != nil {
	//executor.eventService.SetTaskItemStatus([]uint{executor.taskID}, ERROR)
	//executor.NotifyEvent(executor.TaskID, ERROR, "下载文件失败")
	//return err
	//}
	//err = executor.computerService.AddDeviceResource(executor.Option.ID)
	//if err != nil {
	//executor.eventService.SetTaskItemStatus([]uint{executor.taskID}, ERROR)
	//return err
	//}
	//executor.eventService.SetTaskItemStatus([]uint{executor.taskID}, DONE)
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
func (executor *InstallResourceExecutor) BindOption(option interface{}) error {
	output := InstallResourceOption{}
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &output,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(option)
	if err != nil {
		return err
	}
	executor.Option = output
	return nil
}
