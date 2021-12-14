package service

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	model2 "github.com/wenchangshou2/vd-node-manage/module/server/model"
)

type DeviceLayoutOpenService struct {
	ID       uint                   `json:"id" uri:"id"`
	LayoutID string                 `json:"layout_id" binding:"required"`
	Kill     bool                   `json:"kill" binding:"required"`
	Style    map[string]interface{} `json:"style"`
	Windows  []model.Window         `json:"windows"`
}

func (service DeviceLayoutOpenService) Open() serializer.Response {
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
		if window.Source.Category != "id" {
			continue
		}
		for _, v := range ids {
			if window.Source.ID == v {
				continue
			}
		}
		ids = append(ids, window.Source.ID)
	}
	resources, err := model2.GetDeviceResources(service.ID, ids)
	if err != nil {
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
			Style:     window.Style,
			Arguments: window.Arguments,
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
	return serializer.Response{
		Data: c,
	}
}
