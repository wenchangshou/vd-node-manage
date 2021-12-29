package service

import (
	"encoding/json"
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
			"id": id,
		},
	}
}

type DeviceResourceAddService struct {
	ID        uint `json:"id" uri:"id"`
	Resources []struct {
		ID       uint   `json:"ID"`
		Service  string `json:"service"`
		Category string `json:"category" default:"link"`
		Name     string `json:"name"`
		URI      string `json:"uri" binding:"required"`
	} `json:"resources"`
}
type IDRelation struct {
	SId uint `json:"sourceId"`
	DId uint `json:"dstId"`
}

// Add 添加设备资源
func (service DeviceResourceAddService) Add() serializer.Response {
	var (
		rid uint
		err error
	)
	rtu := make([]IDRelation, 0)
	for _, v := range service.Resources {
		resource := model.Resource{
			Name:     v.Name,
			Service:  v.Service,
			Category: v.Category,
			Uri:      v.URI,
		}
		if rid, err = resource.Add(); err != nil {
			return serializer.Err(serializer.CodeDBError, "添加资源失败", err)
		}
		params := make(map[string]interface{})
		params["resource_id"] = rid
		params["uri"] = resource.Uri
		params["name"] = resource.Name
		m := model.Event{
			DeviceID: service.ID,
			Active:   false,
			Action:   model2.InstallResourceAction,
			Status:   model2.Initializes,
		}
		b, _ := json.Marshal(params)
		m.Params = string(b)
		if err = m.Add(); err != nil {
			return serializer.Err(serializer.CodeDBError, "添加资源分发事件失败", err)
		}
		r := IDRelation{SId: v.ID, DId: rid}
		rtu = append(rtu, r)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"relation": rtu,
		},
	}
}
