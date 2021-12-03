package IService

import "github.com/wenchangshou2/vd-node-manage/module/agent-simple/dto"

// TaskService 任务服务接口
type TaskService interface {
	GetTasks() ([]dto.Task, error)
	SetTaskItemStatus([]string, int) error
	SetTaskStatus([]string, int) error
}
