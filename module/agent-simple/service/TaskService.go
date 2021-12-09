package IService

import (
	model2 "github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/dto"
)

// TaskService 任务服务接口
type TaskService interface {
	GetTasks() ([]dto.Task, error)
	SetTaskItemStatus([]uint, int) error
	SetTaskStatus([]uint, model2.EventStatus) error
}
