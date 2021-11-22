package http

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/dto"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	"net/http"
)

type TaskHttpService struct {
	GetUrlStruct
	ID      string
	Address string
}

func (t TaskHttpService) SetTaskItemStatus(ids []string, status int) error {
	rtu := e.HttpBaseData{}
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
func (t TaskHttpService) SetTaskStatus(ids []string, status int) error {
	rtu := e.HttpBaseData{}
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
	Total int        `json:"total"`
	Items []dto.Task `json:"items"`
}
type GetComputerTaskResultForm struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data GetComputerTaskResultDataForm
}

func (t TaskHttpService) GetTasks(status int, count int) ([]dto.Task, error) {
	var rtu GetComputerTaskResultForm
	client := resty.New()
	requestUrl := t.GetUrl(t.Address, fmt.Sprintf("computer/%s/task", t.ID))
	resp, err := client.R().SetBody(map[string]interface{}{
		"status": status,
		"count":  count,
	}).SetResult(&rtu).Get(requestUrl)
	if err != nil {
		return nil, errors.Wrap(err, "请示获取计算机任务失败")
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("获取计算机任务失败,返回的id:%d\n", resp.StatusCode()))
	}
	fmt.Println(rtu)
	return rtu.Data.Items, nil
}

// NewTaskHttpService 创建新的http任务服务
func NewTaskHttpService(id string, address string) TaskHttpService {
	return TaskHttpService{
		ID:   id,
		Address:address,
	}
}
