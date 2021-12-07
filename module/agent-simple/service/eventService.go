package IService

import "github.com/wenchangshou2/vd-node-manage/common/model"

type EventService interface {
	QueryTasks() ([]model.Event, error)
	SetEventStatus([]uint, int) error
}
