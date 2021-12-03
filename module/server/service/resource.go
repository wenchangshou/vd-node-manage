package service

import (
	model2 "github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/model"
)

// ResourceAddService 资源添加服务
type ResourceAddService struct {
	URI      string `json:"uri"`
	Service  string `json:"service"`
	Category string `json:"category"`
	Name     string `json:"name"`
}

func (service *ResourceAddService) Add() serializer.Response {
	var (
		id  uint
		err error
	)
	resource := model.Resource{
		Name:     service.Name,
		Service:  service.Service,
		Category: service.Category,
		Uri:      service.URI,
	}
	if id, err = resource.Add(); err != nil {
		return serializer.Err(serializer.CodeDBError, "添加资源失败", err)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"id":id,
		},
	}
}

type DeviceResourceAddService struct {
	DeviceID   uint `json:"device_id" binding:"required"`
	ResourceID uint `json:"resource_id" binding:"required""`
	Status     int  `json:"status"`
}

// Add 添加设备资源
func (service DeviceResourceAddService) Add() serializer.Response {
	task := model.ResourceDistribution{
		DeviceID:   service.DeviceID,
		ResourceID: service.ResourceID,
		Status:     model2.TaskInitialization,
		Schedule:   0,
	}
	if err := task.Add(); err != nil {
		return serializer.Err(serializer.CodeDBError, "添加设备资源失败", err)
	}
	return serializer.Response{}
}
