package rpc

import (
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/server/g"
	model2 "github.com/wenchangshou/vd-node-manage/module/server/model"
	"github.com/wenchangshou/vd-node-manage/module/server/service"
)

// Register 设备注册
func (device *Device) Register(args *model.DeviceRegisterRequest, reply *model.DeviceRegisterResponse) error {
	service := service.DeviceRegisterService{
		Code:         args.Code,
		ConnType:     args.ConnType,
		HardwareCode: args.HardwareCode,
	}
	id, err := service.Register()
	if err != nil {
		reply.Code = -1
		reply.Msg = err.Error()
	} else {
		reply.HttpAddress = g.Config().Http.Listen
		reply.RpcAddress = g.Config().Listen
		reply.RedisAddress = g.Config().Redis.Addr
		reply.Code = 0
		reply.ID = id
	}
	return nil
}
func (device *Device) SyncConf(_ model.NullRpcRequest, reply *model.DeviceSyncConfResponse) error {
	reply.HttpAddress = g.Config().Http.Listen
	reply.RpcAddress = g.Config().Listen
	reply.RedisAddress = g.Config().Redis.Addr
	reply.Code = 0
	return nil
}
func (device *Device) Ping(_ *model.NullRpcRequest, reply *model.SimpleRpcResponse) error {
	reply.Code = 0
	return nil
}

// ReportStatus 上报状态
func (device *Device) ReportStatus(args *model.DeviceReportRequest, reply *model.SimpleRpcResponse) error {
	if args.ID == 0 {
		reply.Code = 1
		return nil
	}
	cache.Set(fmt.Sprintf("device-%d", args.ID), args.Info, 60)
	//cache.Devices.Put(args)
	return nil
}
func (device *Device) QueryTask(_ *model.DeviceReportRequest, _ *model.DeviceQueryStatusResponse) error {
	return nil
}

// AddDeviceResource 添加设备资源
func (device *Device) AddDeviceResource(args *model.DeviceAddResourceRequest, _ *model.SimpleRpcResponse) error {
	if model2.IsDeviceResource(args.ID, args.ResourceID) {
		return nil
	}

	dr := model2.DeviceResource{
		DeviceID:   args.ID,
		ResourceID: args.ResourceID,
	}
	return dr.Add()
}
func (device Device) DeleteDeviceResource(args *model.DeviceDeleteResourceRequest, _ *model.SimpleRpcResponse) error {
	err := model2.DeleteResourceByDeviceIdAndResourceId(args.ID, args.ResourceID)
	return err
}
