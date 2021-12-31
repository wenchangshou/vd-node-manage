package service

import (
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
)

type DeviceResourceListService struct {
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
	OrderBy    string            `json:"orderBy"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}

type DeviceResourceInfo struct {
	DeviceID   uint   `json:"deviceId"`
	ResourceID uint   `json:"resourceId"`
	Uri        string `json:"uri" gorm:"uri"`
	Name       string `json:"name" gorm:"name"`
	Category   string `json:"category" gorm:"category"`
	Service    string `json:"service" gorm:"service"`
	Status     int    `json:"status" gorm:"status"`
}

func (service DeviceResourceListService) List() serializer.Response {
	var res = make([]DeviceResourceInfo, 0)
	mapping := map[string]string{
		"name":       "resource.name",
		"category":   "resource.category",
		"service":    "resource.service",
		"status":     "resource.status",
		"deviceId":   "device_id",
		"resourceId": "resource_id",
	}
	var total int64 = 0
	tx := model.DB.Table("device_resource")
	tx.Count(&total)
	tx.Select("device_resource.id,device_resource.device_id,device_resource.resource_id," +
		"resource.uri,resource.category,resource.service,resource.status,resource.name")
	if service.OrderBy != "" {
		tx = tx.Order(service.OrderBy)
	}
	for k, v := range service.Conditions {
		if k1, ok := mapping[k]; ok {
			tx = tx.Where(k1+" = ?", v)
		} else {
			tx = tx.Where(k+" = ?", v)
		}
	}
	if service.PageSize > 0 {
		tx = tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize)
	}
	err := tx.Joins("left join resource on resource.id = resource_id  ").Scan(&res).Error
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "调用获取设备资源失败", err)
	}

	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": res,
		},
	}
}

type DeviceResourceDeleteService struct {
	ID         uint `json:"id" uri:"id"`
	ResourceId uint `json:"resourceId" uri:"resource_id"`
}

func (service DeviceResourceDeleteService) Delete() serializer.Response {
	r, err := model.GetDeviceResource(service.ID, service.ResourceId)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "查询资源失败", err)
	}
	m := model.Event{
		DeviceID:   service.ID,
		Active:     false,
		Action:     model2.DeleteResource,
		Status:     model2.Initializes,
		ResourceId: r.ResourceID,
	}
	if err = m.Add(); err != nil {
		return serializer.Err(serializer.CodeDBError, "添加资源事件失败", err)
	}
	return serializer.Response{}
}
