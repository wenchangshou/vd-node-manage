package service

import (
	"errors"
	model2 "github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/model"
	"strings"
)

type DeviceListService struct {
	Page       int               `json:"page" binding:"min=1,required"`
	PageSize   int               `json:"page_size" binding:"min=1,required"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}

func (service *DeviceListService) List() serializer.Response {
	var res []model.Device
	var total int64 = 0
	tx := model.DB.Model(&model.Device{})
	if service.OrderBy != "" {
		tx = tx.Order(service.OrderBy)
	}
	for k, v := range service.Conditions {
		tx = tx.Where(k+" =? ", v)
	}
	if len(service.Searches) > 0 {
		search := ""
		for k, v := range service.Searches {
			search += k + " like '%" + v + "%' OR "
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)

	tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize).Find(&res)
	return serializer.Response{Data: map[string]interface{}{
		"total": total,
		"items": res,
	}}
}

// DeviceRegisterService 设备注册服务
type DeviceRegisterService struct {
	Code         string `json:"code" binding:"required"`
	ConnType     string `json:"connType"`
	HardwareCode string `json:"hardware_code"`
	RpcAddress   string `json:"rpc_address"`
	HttpAddress  string `json:"http_address"`
	RedisAddress string `json:"redis_address"`
}

// Register 注册服务
func (service DeviceRegisterService) Register() (uint, error) {
	device := &model.Device{}
	err := model.DB.Model(&model.Device{}).Where("code=?", service.Code).First(&device).Error
	if err != nil || device.ID <= 0 {
		return 0, errors.New("找不到对应授权id的设备")
	}
	err = model.DB.Debug().Model(device).Updates(map[string]interface{}{
		"status":        model2.Device_Register,
		"conn_type":     service.ConnType,
		"hardware_code": service.HardwareCode,
	}).Error
	if err != nil {
		return 0, errors.New("更新设备状态失败")
	}
	return device.ID, nil
}

type DeviceCreateService struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	ConnType string `json:"conn_type"`
	RegionID int    `json:"region_id"`
}

func (service DeviceCreateService) Create() serializer.Response {
	if model.IsExistsCode(service.Code) {
		return serializer.Err(serializer.CodeDeviceCodeRepeatErr, "授权id已存在", nil)
	}
	device := model.Device{
		Code:     service.Code,
		ConnType: service.ConnType,
		Name:     service.Name,
		Status:   0,
	}
	if service.ConnType == "gateway" {
		device.RegionId = service.RegionID
	}
	err := device.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建设备失败", err)
	}
	return serializer.Response{}
}
func (service DeviceCreateService) List() serializer.Response {

	return serializer.Response{}
}
