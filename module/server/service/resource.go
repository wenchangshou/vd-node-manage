package service

import (
	"context"
	"encoding/json"
	"fmt"
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/g"
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

type ResourceForm struct {
	ID       uint   `json:"id"`
	Service  string `json:"service"`
	Category string `json:"category" default:"link"`
	Name     string `json:"name"`
	Md5      string `json:"md5"`
	URI      string `json:"uri" binding:"required"`
}
type DeviceResourceAddService struct {
	ID        uint           `json:"id" uri:"id"`
	Resources []ResourceForm `json:"resources"`
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
	req, _ := GetEventCmd("checkResourceExists", service.ID, b, true)
	reply, err := g.GEvent.PublishEvent(context.TODO(), fmt.Sprintf("device-%d", service.ID), req, true)
	if err != nil {
		return false
	}
	if err = json.Unmarshal([]byte(reply), &_reply); err != nil {
		return false
	}
	exists, ok := _reply["exists"]
	return ok && exists.(bool)
}
func (service DeviceResourceAddService) add(form ResourceForm) (*model.Resource, error) {
	var (
		_resource *model.Resource
		err       error
	)
	_resource = &model.Resource{
		Name:     form.Name,
		Service:  form.Service,
		Category: form.Category,
		Uri:      form.URI,
		Md5:      form.Md5,
	}
	_, err = _resource.Add()
	return _resource, err
}

// Add 添加设备资源
func (service DeviceResourceAddService) Add() serializer.Response {
	var (
		err error
	)
	rtu := make([]IDRelation, 0)
	for _, v := range service.Resources {
		var _resource *model.Resource
		if v.Service == "web" {
			_resource, err = service.add(v)
			resource := model.DeviceResource{
				DeviceID:   service.ID,
				ResourceID: _resource.ID,
			}
			if err = resource.Add(); err != nil {
				return serializer.Err(serializer.CodeDBError, "添加设备资源失败", err)

			}
			r := IDRelation{SId: v.ID, DId: _resource.ID}
			rtu = append(rtu, r)
			continue
		}
		_resource, err = model.GetResourceByMd5(v.Md5)
		// 如果当前md5的资源不存在，就新建一个资源
		if _resource == nil || err != nil {
			_resource, err = service.add(v)
			if err != nil {
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
