package rpc

import (
	"encoding/json"
	"errors"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	model2 "github.com/wenchangshou2/vd-node-manage/module/server/model"
)

// Query 检测事件
func (event Event) Query(args *model.QueryDeviceEventRequest, reply *model.QueryDeviceEventResponse) error {
	device, err := model2.GetDeviceByID(args.DeviceID)
	if err != nil {
		return errors.New("获取设备失败")
	}
	if device == nil {
		return errors.New("找不到指定的设备")
	}
	events, err := model2.QueryDeviceEventByDeviceID(args.DeviceID)
	if err != nil {
		return errors.New("查询设备事件失败")
	}
	events2 := make([]model.Event, 0)
	for _, event := range events {
		p := make(map[string]interface{})
		json.Unmarshal([]byte(event.Params), &p)
		_e := model.Event{
			ID:       event.ID,
			Name:     event.Name,
			Active:   event.Active,
			DeviceID: event.DeviceID,
			Action:   event.Action,
			Status:   event.Status,
			Params:   p,
		}
		events2 = append(events2, _e)
	}
	reply.Count = len(events2)
	reply.Events = events2
	reply.DeviceID = args.DeviceID
	return nil
}

func (event Event) SetStatus(args *model.DeviceSetEventStatusRequest, reply *model.SimpleRpcResponse) error {

}
