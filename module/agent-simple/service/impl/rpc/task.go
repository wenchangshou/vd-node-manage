package rpc

import (
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
)

type TaskRpcService struct {
	ID uint `json:"id"`
}

func (t TaskRpcService) SetTaskItemStatus(_ []uint, _ int) error {

	panic("implement me")
}

func (t TaskRpcService) SetTaskStatus(ids []uint, status model2.EventStatus) error {

	panic("implement me")
}

// NewTaskRpcService 新的任务rpc服务
func NewTaskRpcService(id uint) *TaskRpcService {
	return &TaskRpcService{
		id,
	}
}
