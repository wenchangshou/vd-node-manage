package executor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/common/file"
	"github.com/wenchangshou/vd-node-manage/common/logging"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
	IService "github.com/wenchangshou/vd-node-manage/module/core/service"
	"github.com/wenchangshou2/zutil"
	bolt "go.etcd.io/bbolt"
	"path"
	"strings"
	"time"
)

type InstallProjectExecutor struct {
	TaskID         uint
	HttpAddress    string
	Options        InstallProjectOption
	NotifyEvent    func(string, int, string)
	HttpRequestUri string
	Mac            string
	eventService   IService.EventService
	DeviceService  IService.DeviceService
	ProjectService IService.ProjectService
	Project        *model.ProjectInfo
	cache          *cache.Driver
	db             *bolt.DB
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

func (executor *InstallProjectExecutor) download(uri, dstPath string) error {
	client := grab.NewClient()
	req, _ := grab.NewRequest(dstPath, uri)
	resp := client.Do(req)
	t := time.NewTicker(time.Second)
loop:
	for {
		select {
		case <-t.C:
			executor.reportDownloadProcess(resp)
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
	logging.GLogger.Info(fmt.Sprintf("eventid:%d download success,save to ./%v", executor.TaskID, resp.Filename))
	return nil
}

//reportDownloadProcess 上报下载进度
func (executor *InstallProjectExecutor) reportDownloadProcess(resp *grab.Response) {
	if executor.cache != nil {
		key := fmt.Sprintf("device-task-%d-%d", g.GetServerInfo().ID, executor.TaskID)
		body := model.TaskDownloadInfo{
			BytesComplete: resp.BytesComplete(),
			Size:          resp.Size,
			Process:       100 * resp.Progress(),
		}
		b, _ := json.Marshal(body)
		(*executor.cache).Set(key, b, 24*60*60)
	}

}

func (executor *InstallProjectExecutor) Execute() error {
	var (
		err        error
		_md5       string
		targetPath string
	)

	cfg := g.Config()
	dstPath := path.Join(cfg.Resource.Tmp)
	dstPath = path.Join(dstPath, fmt.Sprintf("%d-%s", executor.Project.ID, executor.Project.Name))
	zutil.IsExistDelete(dstPath)
	if executor.download(executor.Project.Uri, dstPath); err != nil {
		err = fmt.Errorf("%s:%v", "下载文件失败", err)
		goto End
	}
	if _md5, err = zutil.GetFileMd5(dstPath); err != nil {
		logging.GLogger.Warn(fmt.Sprintf("校验文件md5失败,失败原因:%s", err.Error()))
		err = errors.New("计算文件md5错误:" + err.Error())
		goto End
	}
	if !strings.EqualFold(_md5, executor.Project.Md5) {
		err = fmt.Errorf("md5错误,期望的md5:%s,实际文件的md5:%s", executor.Project.Md5, _md5)
		goto End
	}
	targetPath = path.Join(cfg.Resource.Directory, "project/", fmt.Sprintf("%d-%s", executor.Project.ID, executor.Project.Name))
	zutil.IsNotExistMkDir(targetPath)
	_, err = file.DeCompress(dstPath, targetPath)
	if err != nil {
		err = fmt.Errorf("解压zip失败:%v", err)
		goto End
	}
	if err = executor.DeviceService.AddComputerProject(executor.Project.ID); err != nil {
		return errors.New("添加设备项目失败:" + err.Error())
	}
End:
	zutil.IsExistDelete(dstPath)
	if err == nil {
		executor.storeProjectData()
	}
	return err
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
func (executor *InstallProjectExecutor) BindOption(option interface{}) error {
	return nil
}
func (executor *InstallProjectExecutor) SubscribeNotifyStatusChange(event func(string, int, string)) {
	executor.NotifyEvent = event

}

func (executor *InstallProjectExecutor) storeProjectData() {
	executor.db.Update(func(tx *bolt.Tx) error {
		var (
			bucket *bolt.Bucket
			err    error
		)
		key := fmt.Sprintf("project-%d", executor.Project.ID)
		bucket, err = tx.CreateBucketIfNotExists([]byte(key))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("md5"), []byte(executor.Project.Md5))
		if err != nil {
			return err
		}
		return nil
	})
}
