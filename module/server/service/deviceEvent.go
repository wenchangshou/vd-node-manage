package service

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
	"strconv"
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
	ID         uint                    `json:"id"`
	Name       string                  `json:"name"`
	Active     int                     `json:"active"`
	DeviceId   uint                    `json:"deviceId"`
	ResourceId uint                    `json:"resourceId"`
	Action     int                     `json:"action"`
	Status     int                     `json:"status"`
	Download   model2.TaskDownloadInfo `json:"download"  gorm:"-" default:"{}"`
}

func (service DeviceEventGetService) Get() serializer.Response {
	deviceId := ""
	kc := map[string]string{
		"deviceId":   "device_id",
		"resourceId": "resource_id",
	}
	ids := make([]string, 0)
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
		if k == "deviceId" {
			deviceId = v
		}
		if k1, ok := kc[k]; ok {
			tx = tx.Where(k1+" = ?", v)
		} else {
			tx = tx.Where(k+" = ?", v)
		}
	}
	if deviceId == "" {
		return serializer.Err(serializer.CodeParamErr, "deviceId必須存在", nil)
	}

	if service.PageSize > 0 {
		tx = tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize)
	}
	err := tx.Find(&res).Error
	if err != nil {
		serializer.Err(serializer.CodeDBError, "获取事件失败", err)
	}
	for _, v := range res {
		_id := strconv.Itoa(int(v.ID))
		ids = append(ids, _id)
	}
	m, _ := cache.GetSettings(ids, fmt.Sprintf("device-task-%s-", deviceId))
	for k, v := range res {
		//k2 := fmt.Sprintf("device-task-%d-%d", v.DeviceId, v.ResourceId)
		_id := strconv.Itoa(int(v.ID))
		if b, exists := m[_id]; exists {
			t := model2.TaskDownloadInfo{}
			json.Unmarshal([]byte(b), &t)
			res[k].Download = t
		}
	}
	return serializer.Response{Data: map[string]interface{}{
		"total": total,
		"items": res,
	}}
}
