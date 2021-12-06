package rpc

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	model2 "github.com/wenchangshou2/vd-node-manage/module/server/model"
)

//func (task Task) QueryDeviceResourceDistribution(args *model.QueryDeviceResourceDistributionRequest, reply *model.QueryDeviceResourceDistributionResponse) error {
//	device, err := model2.GetDeviceByID(args.DeviceID)
//	if err != nil {
//		return errors.New("获取设备失败")
//	}
//	if device == nil {
//		return errors.New("找不到指定的设备")
//	}
//
//	tasks, err := model2.QueryResourceDistributionByDeviceID(args.DeviceID)
//	if err != nil {
//		return err
//	}
//	reply.Count = len(tasks)
//	reply.Tasks = tasks
//	reply.DeviceID = args.DeviceID
//	return nil
//}
//

func (task Task) SetResourceDistributionStatus(args *model.DeviceSetStatusRequest, reply *model.SimpleRpcResponse) error {
	err := model2.SetResourceDistributionStatus(args.ID, args.Status)
	if err != nil {
		reply.Code = 10000
		return err
	}
	return nil
}
