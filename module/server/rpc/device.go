package rpc

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/cache"
	"github.com/wenchangshou2/vd-node-manage/module/server/service"
)

func (device *Device) Register(args *model.DeviceRegisterRequest, reply *model.DeviceRegisterResponse) error {
	service := service.DeviceRegisterService{
		Code:     args.Code,
		ConnType: args.ConnType,
	}
	id, err := service.Register()
	if err != nil {
		reply.Code = -1
		reply.Msg = err.Error()
	} else {
		reply.Code = 0
		reply.ID = id
	}
	return nil
}

// ReportStatus 上报状态
func (device *Device) ReportStatus(args *model.DeviceReportRequest, reply *model.SimpleRpcResponse) error {
	if args.ID == "" {
		reply.Code = 1
		return nil
	}
	cache.Devices.Put(args)
	return nil
}
func (device *Device) QueryTask() error {
	return nil
}
