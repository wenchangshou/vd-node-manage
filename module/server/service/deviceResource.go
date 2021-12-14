package service

import (
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/model"
)

type DeviceResourceListService struct {
	ID uint `json:"id" uri:"id"`
}

func (service DeviceResourceListService) List() serializer.Response {
	devices, err := model.GetDeviceResourcesByDeviceID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取设备资源列表失败", err)
	}
	return serializer.Response{
		Data: devices,
	}
}
