package service

import "github.com/wenchangshou2/vd-node-manage/common/serializer"

type DeviceResourceListService struct {
	ID uint `json:"id" uri:"id"`
}

func (service DeviceResourceListService) List() serializer.Response {
	return serializer.Response{}
}
