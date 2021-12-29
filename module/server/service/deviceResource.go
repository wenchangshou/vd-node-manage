package service

import (
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/model"
)

type DeviceResourceListService struct {
	ID       uint `json:"id" uri:"id"`
	PageSize int  `json:"page_size" form:"page_size"`
	Page     int  `json:"page" form:"page"`
}

type DeviceResourceInfo struct {
	ID       uint   `json:"ID"`
	Name     string `json:"name"`
	Service  string `json:"service"`
	Category string `json:"category"`
	Status   uint   `json:"status"`
}

func (service DeviceResourceListService) List() serializer.Response {
	var res = make([]DeviceResourceInfo, 0)
	var tmp []model.DeviceResource
	var total int64 = 0
	tx := model.DB.Model(&model.DeviceResource{})
	tx.Where("device_id=?", service.ID)
	tx.Preload("Resource")
	tx.Count(&total)
	err := tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize).Find(&tmp).Error
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取设备资源列表失败", err)
	}
	for _, v := range tmp {
		t := DeviceResourceInfo{
			ID:       v.Resource.ID,
			Name:     v.Resource.Name,
			Service:  v.Resource.Service,
			Category: v.Resource.Category,
			Status:   v.Resource.Status,
		}
		res = append(res, t)
	}

	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": res,
		},
	}
}
