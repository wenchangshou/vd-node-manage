package rpc

import (
	"errors"

	"github.com/wenchangshou/vd-node-manage/common/model"
	model2 "github.com/wenchangshou/vd-node-manage/module/server/model"
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
	events, err := model2.QueryDeviceEventByDeviceIDAndStatus(args.DeviceID, args.Status)
	if err != nil {
		return errors.New("查询设备事件失败")
	}
	events2 := make([]model.Event, 0)
	for _, event := range events {

		_e := model.Event{
			ID:         event.ID,
			Name:       event.Name,
			Active:     event.Active,
			DeviceID:   event.DeviceID,
			Action:     event.Action,
			Status:     event.Status,
			ResourceId: event.ResourceId,
			ProjectId:  event.ProjectId,
		}
		events2 = append(events2, _e)
	}
	reply.Count = len(events2)
	reply.Events = events2
	reply.DeviceID = args.DeviceID
	return nil
}

func (event Event) SetStatus(args *model.DeviceSetEventStatusRequest, _ *model.SimpleRpcResponse) error {
	for _, id := range args.EventID {
		err := model2.SetDeviceEventStatus(id, args.Status)
		if err != nil {
			return err
		}

	}
	return nil
}
