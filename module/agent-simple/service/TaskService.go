package IService

import (
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
)

// TaskService 任务服务接口
type TaskService interface {
	SetTaskItemStatus([]uint, int) error
	SetTaskStatus([]uint, model2.EventStatus) error
}
