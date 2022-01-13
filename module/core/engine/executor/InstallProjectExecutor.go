package executor

import (
	"encoding/json"
)

type InstallProjectExecutor struct {
	TaskID         uint
	HttpAddress    string
	Options        InstallProjectOption
	NotifyEvent    func(string, int, string)
	HttpRequestUri string
	Mac            string
}
type File struct {
	Name       string `gorm:"name"`
	Mode       string `gorm:"mode"`
	SourceName string `gorm:"source_name"`
	UserId     string `gorm:"user_id"`
	Size       uint   `gorm:"size"`
	Uuid       string `gorm:"uuid"`
}

func (file File) GetResourcePath() string {
	return file.Uuid + file.SourceName
}
func (file *File) GetApplicationPath() string {
	return file.Uuid
}

type InstallProjectOption struct {
	ID uint `json:"id"`
	//ProjectReleaseID string `json:"project_release_id"`
	Uri    string `json:"uri"`
	Source string `json:"source"`
	Name   string `json:"name"`
}

func (executor *InstallProjectExecutor) Execute() error {
	//var (
	//	err error
	//)
	//cfg := g.Config()
	////requestUri := "http://" + executor.HttpRequestUri + "/" + executor.Options.Uri
	//dstPath := path.Join(cfg.Resource.Directory, "application", string(executor.Options.ID))
	//tmpPath := path.Join(cfg.Resource.Tmp, executor.Options.Source)
	//err = file.DownloadFile(executor.Options.Uri, cfg.Resource.Tmp, executor.Options.Source, func(length, downLen int64) {
	//
	//})
	//if err != nil {
	//	fmt.Println("下载错误:", err)
	//}
	//zutil.IsNotExistMkDir(dstPath)
	//err = util.UnZip(dstPath, tmpPath)
	//if err != nil {
	//	os.RemoveAll(tmpPath)
	//	executor.TaskService.SetTaskItemStatus([]uint{executor.TaskID}, ERROR)
	//	return err
	//}
	//err = executor.ComputerService.AddComputerProject(executor.Options.ID)
	//if err != nil {
	//	executor.TaskService.SetTaskItemStatus([]uint{executor.TaskID}, ERROR)
	//	return err
	//}
	//os.RemoveAll(tmpPath)
	//executor.TaskService.SetTaskItemStatus([]uint{executor.TaskID}, DONE)
	return nil
}

// Cancel 远程任务取消
func (executor *InstallProjectExecutor) Cancel() error {
	return nil
}

// Verification 检验任务参数
func (executor *InstallProjectExecutor) Verification(option string) bool {
	err := json.Unmarshal([]byte(option), &executor.Options)
	return err == nil
}

// BindOption  检验任务参数
func (executor *InstallProjectExecutor) BindOption(option string) error {
	err := json.Unmarshal([]byte(option), &executor.Options)
	return err
}
func (executor *InstallProjectExecutor) SubscribeNotifyStatusChange(event func(string, int, string)) {
	executor.NotifyEvent = event

}