package rpc

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/server/g"
	"github.com/wenchangshou2/vd-node-manage/module/server/service"
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
func (device *Device) Ping(args *model.NullRpcRequest, reply *model.SimpleRpcResponse) error {
	reply.Code = 0
	return nil
}

// ReportStatus 上报状态
func (device *Device) ReportStatus(args *model.DeviceReportRequest, reply *model.SimpleRpcResponse) error {
	if args.ID == "" {
		reply.Code = 1
		return nil
	}
	//cache.Devices.Put(args)
	return nil
}
func (device *Device) QueryTask(args *model.DeviceReportRequest, reply *model.DeviceQueryStatusResponse) error {
	return nil
}
