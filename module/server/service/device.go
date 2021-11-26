package service

import (
	"errors"
	model2 "github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/model"
)

// DeviceRegisterService 设备注册服务
type DeviceRegisterService struct {
	Code     string `json:"code" binding:"required"`
	ConnType string `json:"connType"`
}

// Register 注册服务
func (service DeviceRegisterService) Register() (uint,error) {
	device := &model.Device{}
	err := model.DB.Model(&model.Device{}).Where("code=?", service.Code).First(&device).Error
	if err != nil || device.ID <= 0 {
		return 0,errors.New("找不到对应授权id的设备")
	}
	err = model.DB.Debug().Model(device).Updates(map[string]interface{}{
		"status":   model2.Device_Init,
		"connType": service.ConnType,
	}).Error
	if err != nil {
		return 0,errors.New("更新设备状态失败")
	}
	return device.ID,nil
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
