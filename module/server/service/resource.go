package service

import (
	"context"
	"encoding/json"
	"fmt"
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/event"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
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
		Md5      string `json:"md5"`
		URI      string `json:"uri" binding:"required"`
	} `json:"resources"`
}
type IDRelation struct {
	SId uint `json:"sourceId"`
	DId uint `json:"dstId"`
}

func (service DeviceResourceAddService) CheckRemoteDeviceResource(deviceId uint, r *model.Resource) bool {
	_reply := make(map[string]interface{})
	_m := make(map[string]interface{})
	_m["id"] = r.ID
	_m["name"] = r.Name
	_m["md5"] = r.Md5
	b, _ := json.Marshal(_m)
	dr, err := model.GetDeviceResource(deviceId, r.ID)
	if dr == nil || dr.ResourceID <= 0 {
		return false
	}
	req, _ := event.GetEventCmd("checkResourceExists", service.ID, b, true)
	reply, err := event.GEvent.PublishEvent(context.TODO(), fmt.Sprintf("device-%d", service.ID), req, true)
	if err != nil {
		return false
	}
	if err = json.Unmarshal([]byte(reply), &_reply); err != nil {
		return false
	}
	exists, ok := _reply["exists"]
	return ok && exists.(bool)
}

// Add 添加设备资源
func (service DeviceResourceAddService) Add() serializer.Response {
	var (
		err       error
		_resource *model.Resource
	)
	rtu := make([]IDRelation, 0)
	for _, v := range service.Resources {
		_resource, err = model.GetResourceByMd5(v.Md5)
		// 如果当前md5的资源不存在，就新建一个资源
		if _resource == nil || err != nil {
			_resource = &model.Resource{
				Name:     v.Name,
				Service:  v.Service,
				Category: v.Category,
				Uri:      v.URI,
				Md5:      v.Md5,
			}
			if _, err := _resource.Add(); err != nil {
				return serializer.Err(serializer.CodeDBError, "添加资源失败", err)
			}
		}
		m := model.Event{
			DeviceID: service.ID,
			Active:   false,
			Action:   model2.InstallResourceAction,
			Status:   model2.Initializes,
		}
		// 如果远端已经存在了对应的资源，就不走下发
		//if service.CheckRemoteDeviceResource(service.ID, _resource) {
		//	m.Status = model2.Done
		//}
		m.ResourceId = _resource.ID
		if err = m.Add(); err != nil {
			return serializer.Err(serializer.CodeDBError, "添加资源分发事件失败", err)
		}
		r := IDRelation{SId: v.ID, DId: _resource.ID}
		rtu = append(rtu, r)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"relation": rtu,
		},
	}
}
