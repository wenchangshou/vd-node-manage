package computer

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ListComputerTaskService struct {
	ID         string            `json:"id" uri:"id" form:"id"`
	Page       int               `json:"page" `
	PageSize   int               `json:"page_size"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
	Status     int               `json:"status" form:"status" uri:"status"`
	Count      int               `json:"count" form:"count" uri:"count"`
}

func (service ListComputerTaskService) GetComputerTask() serializer.Response {
	var (
		res []model.Task
		err error
	)
	computer, err := model.GetComputerById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机对象失败", err)
	}
	res, err = computer.GetTasks(model.TaskStatus(service.Status), service.PageSize)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取任务失败", err)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"items": res,
			"total": len(res),
		},
	}
}
