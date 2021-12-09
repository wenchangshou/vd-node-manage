package IService

import "github.com/wenchangshou2/vd-node-manage/common/model"

type EventService interface {
	QueryDeviceEvent(status model.EventStatus) ([]model.Event, error)
	SetEventStatus([]uint, model.EventStatus) error
}
