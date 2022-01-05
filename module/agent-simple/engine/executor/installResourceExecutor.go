package executor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/mitchellh/mapstructure"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
	IService "github.com/wenchangshou/vd-node-manage/module/agent-simple/service"
	"github.com/wenchangshou2/zutil"
	bolt "go.etcd.io/bbolt"
	"path"
	"strings"
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
	cache           *cache.Driver
	db              *bolt.DB
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
	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	return nil
}
func (executor *InstallResourceExecutor) localFileExists(rPath string) bool {
	if !zutil.IsExist(rPath) {
		return false
	}
	exists := true
	err := executor.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(fmt.Sprintf("resource-%d", executor.Resource.ID)))
		if b == nil {
			exists = false
			return nil
		}
		_md5 := b.Get([]byte("md5"))
		if _md5 == nil || string(_md5) != executor.Resource.Md5 {
			exists = false
			return nil
		}
		return nil
	})
	return err == nil && exists

}
func (executor *InstallResourceExecutor) storeResourceData() {
	executor.db.Update(func(tx *bolt.Tx) error {
		var (
			bucket *bolt.Bucket
			err    error
		)
		key := fmt.Sprintf("resource-%d", executor.Resource.ID)
		bucket, err = tx.CreateBucketIfNotExists([]byte(key))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("md5"), []byte(executor.Resource.Md5))
		if err != nil {
			return err
		}
		return nil
	})
}
func (executor *InstallResourceExecutor) Execute() error {
	cfg := g.Config()
	dstPath := path.Join(cfg.Resource.Directory, "resource/")
	dstPath = path.Join(dstPath, fmt.Sprintf("%d-%s", executor.Resource.ID, executor.Resource.Name))
	exists := executor.localFileExists(dstPath)
	// 防止重复下载
	if exists {
		return nil
	}
	zutil.IsExistDelete(dstPath)
	err := executor.download(executor.Resource.Uri, dstPath)
	if err != nil {
		return fmt.Errorf("%s:%v", "下载文件失败", err)
	}
	_md5, err := zutil.GetFileMd5(dstPath)
	if err != nil {
		return errors.New("计算文件md5错误:" + err.Error())
	}
	if strings.ToLower(_md5) != strings.ToLower(executor.Resource.Md5) {
		return fmt.Errorf("md5错误,期望的md5:%s,实际文件的md5:%s", executor.Resource.Md5, _md5)
	}
	if err = executor.DeviceService.AddComputerResource(executor.Resource.ID); err != nil {
		return errors.New("添加设备资源失败:" + err.Error())
	}
	executor.storeResourceData()
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

//reportDownloadProcess 上报下载进度
func (executor *InstallResourceExecutor) reportDownloadProcess(resp *grab.Response) {
	if executor.cache != nil {
		key := fmt.Sprintf("device-task-%d-%d", g.Config().Server.ID, executor.taskID)
		body := model.TaskDownloadInfo{
			BytesComplete: resp.BytesComplete(),
			Size:          resp.Size,
			Process:       100 * resp.Progress(),
		}
		b, _ := json.Marshal(body)
		(*executor.cache).Set(key, b, 24*60*60)
	}

}
