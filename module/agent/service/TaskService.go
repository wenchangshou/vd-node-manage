package IService

import "github.com/wenchangshou2/vd-node-manage/module/agent/dto"

// TaskService 任务服务接口
type TaskService interface {
	GetTasks(status int, count int) ([]dto.Task, error)
	SetTaskItemStatus([]string, int) error
	SetTaskStatus([]string, int) error
}
