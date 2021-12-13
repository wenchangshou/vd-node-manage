package rpc

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
)

type EventRpcService struct {
	Client *g.SingleConnRpcClient
	ID     uint `json:"id"`
}

func (event EventRpcService) QueryDeviceEvent(status model.EventStatus) ([]model.Event, error) {
	req := model.QueryDeviceEventRequest{DeviceID: event.ID, Status: status}
	reply := model.QueryDeviceEventResponse{}
	err := event.Client.Call("Event.Query", &req, &reply)
	return reply.Events, err
}
func (event EventRpcService) SetEventStatus(id []uint, status model.EventStatus) error {
	req := model.DeviceSetEventStatusRequest{
		EventID: id,
		Status:  status,
	}
	reply := model.SimpleRpcResponse{}
	return event.Client.Call("Event.SetStatus", &req, &reply)
}

func NewEventRpcService(id uint, client *g.SingleConnRpcClient) *EventRpcService {

	return &EventRpcService{ID: id, Client: client}
}
