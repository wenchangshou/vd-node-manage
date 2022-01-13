package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"net/http"
)

type TaskHttpService struct {
	GetUrlStruct
	ID      string
	Address string
}

func (t TaskHttpService) SetTaskItemStatus(ids []uint, status int) error {
	rtu := model.HttpBaseData{}
	client := resty.New()
	requestUri := t.GetUrl(t.Address, "task")
	resp, err := client.R().SetBody(map[string]interface{}{
		"id":     ids,
		"status": status,
	}).SetResult(&rtu).Post(requestUri)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return errors.Wrap(err, "调用更改任务状态接口失败")
	}
	if rtu.Code != 0 {
		return errors.New("调用更改任务状态接口失败:" + rtu.Msg)
	}
	return nil
}

// SetTaskStatus 设置任务状态
func (t TaskHttpService) SetTaskStatus(ids []uint, status int) error {
	rtu := model.HttpBaseData{}
	client := resty.New()
	requestUri := t.GetUrl(t.Address, "task")
	resp, err := client.R().SetBody(map[string]interface{}{
		"id":     ids,
		"status": status,
	}).SetResult(&rtu).Post(requestUri)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return errors.Wrap(err, "调用更改子任务状态接口失败")
	}
	if rtu.Code != 0 {
		return errors.New("调用更改子任务状态接口失败:" + rtu.Msg)
	}
	return nil
}

type GetComputerTaskResultDataForm struct {
	Total int `json:"total"`
}
type GetComputerTaskResultForm struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data GetComputerTaskResultDataForm
}

// NewTaskHttpService 创建新的http任务服务
func NewTaskHttpService(id string, address string) TaskHttpService {
	return TaskHttpService{
		ID:      id,
		Address: address,
	}
}
