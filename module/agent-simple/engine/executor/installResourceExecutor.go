package executor

import (
	"errors"
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/mitchellh/mapstructure"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
	IService "github.com/wenchangshou/vd-node-manage/module/agent-simple/service"
	"github.com/wenchangshou2/zutil"
	"path"
	"time"
)

type InstallResourceExecutor struct {
	HttpRequestUri  string
	eventService    IService.EventService
	DeviceService   IService.DeviceService
	Mac             string
	taskID          uint
	ResourceService IService.ResourceService
	Resource        *model.ResourceInfo
}
type InstallResourceOption struct {
	ResourceID uint   `json:"resource_id"`
	Name       string `json:"name"`
	Uri        string `json:"uri"`
}

func (executor *InstallResourceExecutor) download(uri, dstPath string) error {
	client := grab.NewClient()
	req, _ := grab.NewRequest(dstPath, uri)
	resp := client.Do(req)
	t := time.NewTicker(500 * time.Millisecond)
loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())
		case <-resp.Done:
			break loop
		}
	}
	if err := resp.Err(); err != nil {
		return err
	}
	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	return nil
}
func (executor *InstallResourceExecutor) Execute() error {
	cfg := g.Config()

	dstPath := path.Join(cfg.Resource.Directory, "resource/")
	dstPath = path.Join(dstPath, fmt.Sprintf("%d-%s", executor.Resource.ID, executor.Resource.Name))
	zutil.IsExistDelete(dstPath)
	err := executor.download(executor.Resource.Uri, dstPath)
	if err != nil {
		return fmt.Errorf("%s:%v", "下载文件失败", err)
	}
	if err = executor.DeviceService.AddComputerResource(executor.Resource.ID); err != nil {
		return errors.New("添加设备资源失败:" + err.Error())
	}
	return nil
}

func (executor *InstallResourceExecutor) Cancel() error {
	return nil
}
func (executor *InstallResourceExecutor) Verification(option string) bool {
	return true
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
	return nil
}
