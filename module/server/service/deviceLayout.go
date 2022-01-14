package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/g"
	model2 "github.com/wenchangshou/vd-node-manage/module/server/model"
)

func GetEventCmd(action string, did uint, body []byte, reply bool) (*model.EventRequest, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	msg := model.EventRequest{
		EventID:   id.String(),
		Action:    action,
		DeviceID:  did,
		Arguments: body,
		Reply:     reply,
	}
	return &msg, nil

}

type DeviceLayoutOpenWindowService struct {
	ID         uint   `json:"id" uri:"id"`
	LayoutId   string `json:"layoutId" uri:"layoutId"`
	WindowId   string `json:"windowId" uri:"windowId"`
	ResourceId uint   `json:"resourceId"`
}

func (service DeviceLayoutOpenWindowService) Open() serializer.Response {
	deviceResource, err := model2.GetDeviceResource(service.ID, service.ResourceId)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取资源失败", err)
	}
	req := model.OpenWindowCmdParams{
		ID:       service.ID,
		LayoutID: service.LayoutId,
		WindowID: service.WindowId,
		Service:  deviceResource.Resource.Service,
		Source:   fmt.Sprintf("%d-%s", deviceResource.ResourceID, deviceResource.Resource.Name),
	}
	b1, _ := json.Marshal(req)
	r, _ := GetEventCmd("change", service.ID, b1, true)
	reply, err := g.GEvent.PublishEvent(context.TODO(), fmt.Sprintf("device-%d", service.ID), r, true)
	if err != nil {
		return serializer.Err(serializer.CodeRedisError, "publish event error", err)
	}
	return serializer.Response{Data: reply}
}

// DeviceLayoutOpenService 打开设备布局服务
type DeviceLayoutOpenService struct {
	ID       uint                   `json:"id" uri:"id"`
	Startup  bool                   `json:"startup"`
	LayoutID string                 `uri:"layout_id" json:"layout_id"`
	Kill     bool                   `json:"kill" binding:"required"`
	Style    map[string]interface{} `json:"style"`
	Windows  []model.Window         `json:"windows"`
}

func (service DeviceLayoutOpenService) Open() serializer.Response {
	var (
		resources []model2.DeviceResource
		err       error
	)
	c := model.OpenLayoutCmdParams{
		ID:    service.LayoutID,
		Kill:  service.Kill,
		Style: make(map[string]interface{}),
	}
	if service.Style != nil {
		c.Style = service.Style
	}
	windows := make([]model.OpenWindowInfo, 0)
	mapDeviceResource := make(map[uint]model2.DeviceResource)
	ids := make([]uint, 0)
	for _, window := range service.Windows {
		flag := true
		if window.Source.Category != "id" {
			continue
		}
		for _, v := range ids {
			if window.Source.ID == v {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		ids = append(ids, window.Source.ID)
	}
	if resources, err = model2.GetDeviceResources(service.ID, ids); err != nil {
		return serializer.Err(serializer.CodeDBError, "获取设备资源失败", err)
	}
	for _, resource := range resources {
		mapDeviceResource[resource.ResourceID] = resource
	}
	for _, window := range service.Windows {
		w := model.OpenWindowInfo{
			ID:        window.ID,
			X:         window.X,
			Y:         window.Y,
			Z:         window.Z,
			Width:     window.Width,
			Height:    window.Height,
			Service:   window.Service,
			Style:     make(map[string]interface{}),
			Arguments: make(map[string]interface{}),
		}
		if window.Style != nil {
			w.Style = window.Style
		}
		if window.Arguments != nil {
			w.Arguments = window.Arguments
		}
		r := mapDeviceResource[window.Source.ID]
		if window.Service == "http" {
			w.Source = r.Resource.Uri
		} else {
			w.Source = fmt.Sprintf("%d-%s", r.ResourceID, r.Resource.Name)
		}
		windows = append(windows, w)
	}
	c.Windows = windows
	b1, _ := json.Marshal(c)
	ctx := context.TODO()
	r, _ := GetEventCmd("openLayout", service.ID, b1, true)
	reply, err := g.GEvent.PublishEvent(ctx, fmt.Sprintf("device-%d", service.ID), r, true)
	if err != nil {
		return serializer.Err(serializer.CodeRedisError, "redis publish event error", err)
	}
	if service.Startup {
		err := model2.SetDeviceStartup(service.ID, b1)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "设置默认启动失败", err)
		}
	}
	return serializer.Response{
		Data: reply,
	}
}

type DeviceLayoutCloseService struct {
	ID uint `json:"id" uri:"id"`
}

func (service DeviceLayoutCloseService) Close() serializer.Response {
	r, _ := GetEventCmd("closeLayout", service.ID, nil, true)
	reply, err := g.GEvent.PublishEvent(context.TODO(), fmt.Sprintf("device-%d", service.ID), r, true)
	if err != nil {
		return serializer.Err(serializer.CodeRedisError, "redis publish event error", err)
	}
	return serializer.Response{
		Data: reply,
	}
}

type DeviceLayoutControlService struct {
	ID   uint   `json:"id" uri:"id"`
	Lid  string `json:"layout_id" uri:"layout_id"`
	Wid  string `json:"window_id" uri:"window_id"`
	Body string `json:"body"`
}

func (service DeviceLayoutControlService) Control() serializer.Response {
	c := model.ControlWindowCmdParams{
		ID:   service.Lid,
		Wid:  service.Wid,
		Body: service.Body,
	}
	b1, _ := json.Marshal(c)
	r, _ := GetEventCmd("control", service.ID, b1, true)
	reply, err := g.GEvent.PublishEvent(context.TODO(), fmt.Sprintf("device-%d", service.ID), r, true)
	if err != nil {
		return serializer.Err(serializer.CodeRedisError, "redis publish event error", err)
	}
	return serializer.Response{Data: reply}
}

type DeviceLayoutGetService struct {
	ID       int    `json:"id" uri:"id"`
	LayoutID string `json:"layout_id" uri:"layout_id"`
	Wid      string `json:"wid" uri:"window_id"`
}

func (service DeviceLayoutGetService) Get() serializer.Response {
	val, exists := cache.Get(fmt.Sprintf("device-%d-%s", service.ID, service.LayoutID))
	obj := make([]model.ActiveWindowInfo, 0)
	if !exists || val == nil {
		return serializer.Err(serializer.CodeGetLayoutInfoFail, "没有找到指定的布局信息", nil)
	}
	err := json.Unmarshal([]byte(val.(string)), &obj)
	if err != nil {
		return serializer.Err(serializer.CodeGetLayoutInfoFail, "实时数据异常", err)
	}
	return serializer.Response{
		Data: obj,
	}
}
func (service DeviceLayoutGetService) GetWindow() serializer.Response {
	val, exists := cache.Get(fmt.Sprintf("device-%d-%s", service.ID, service.LayoutID))
	obj := make([]model.ActiveWindowInfo, 0)
	if !exists || val == nil {
		return serializer.Err(serializer.CodeGetLayoutInfoFail, "没有找到指定的布局信息", nil)
	}
	err := json.Unmarshal([]byte(val.(string)), &obj)
	if err != nil {
		return serializer.Err(serializer.CodeGetLayoutInfoFail, "实时数据异常", err)
	}
	for _, v := range obj {
		if v.ID == service.Wid {
			return serializer.Response{
				Data: v,
			}
		}
	}
	return serializer.Err(serializer.CodeGetLayoutInfoFail, "没有找到指定的容器", nil)
}
