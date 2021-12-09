package rpc

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"time"
)

type EventRpcService struct {
	ID uint `json:"id"`
}

func (event EventRpcService) QueryDeviceEvent(status model.EventStatus) ([]model.Event, error) {
	rpcClient := &g.SingleConnRpcClient{
		RpcServer: fmt.Sprintf(g.Config().Server.RpcAddress),
		Timeout:   time.Second,
	}
	req := model.QueryDeviceEventRequest{DeviceID: event.ID, Status: status}
	reply := model.QueryDeviceEventResponse{}
	err := rpcClient.Call("Event.Query", &req, &reply)
	return reply.Events, err
}
func (event EventRpcService) SetEventStatus(id []uint, status model.EventStatus) error {
	c := &g.SingleConnRpcClient{
		RpcServer: fmt.Sprintf(g.Config().Server.RpcAddress),
		Timeout:   time.Second,
	}
	req := model.DeviceSetStatusRequest{
		ID:     id,
		Status: status,
	}
	reply := model.SimpleRpcResponse{}
	return c.Call("Event.SetStatus", &req, &reply)
}

func NewEventRpcService(id uint) *EventRpcService {
	return &EventRpcService{ID: id}
}
