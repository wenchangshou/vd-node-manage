package task

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type QueryComputerTaskService struct {
	Mac string `json:"mac" uri:"mac" form:"mac"`
}

// Query 查询计算机任务
func (service *QueryComputerTaskService) Query() serializer.Response {
	computer, err := model.GetComputerByMac(service.Mac)
	if err != nil {
		return serializer.Err(serializer.CodeNoFoundComputerErr, "没有找到指定计算机", nil)
	}
	tasks, err := model.GetPendingTaskByComputerId(computer.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取待处理任务失败", err)
	}
	return serializer.Response{
		Data: tasks,
	}
}
