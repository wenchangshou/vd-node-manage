package service

import (
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
	"strings"
)

type DeviceEventGetService struct {
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
	OrderBy    string            `json:"orderBy"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}
type EventInfo struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Active     int    `json:"active"`
	DeviceId   uint   `json:"deviceId"`
	ResourceId uint   `json:"resourceId"`
	Action     int    `json:"action"`
	Status     int    `json:"status"`
}

func (service DeviceEventGetService) Get() serializer.Response {
	kc := map[string]string{
		"deviceId":   "device_id",
		"resourceId": "resource_id",
	}
	var res []EventInfo
	var total int64 = 0
	tx := model.DB.Model(&model.Event{})
	tx.Count(&total)
	for k, v := range service.Conditions {
		if k == "resourceId" {
			arr := strings.Split(v, ",")
			tx = tx.Where("resource_id IN ?", arr)
			continue
		}
		if k1, ok := kc[k]; ok {
			tx = tx.Where(k1+" = ?", v)
		} else {
			tx = tx.Where(k+" = ?", v)
		}
	}
	if service.PageSize > 0 {
		tx = tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize)
	}
	err := tx.Find(&res).Error
	if err != nil {
		serializer.Err(serializer.CodeDBError, "获取事件失败", err)
	}
	return serializer.Response{Data: map[string]interface{}{
		"total": total,
		"items": res,
	}}
}
