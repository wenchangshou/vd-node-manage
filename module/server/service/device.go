package service

import (
	"errors"
	"github.com/google/uuid"
	model2 "github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/model"
	"github.com/wenchangshou2/vd-node-manage/module/server/vo"
)

type DeviceListService struct {
	Page     int `json:"page" binding:"min=1,required" form:"page"`
	PageSize int `json:"pageSize" binding:"min=1,required" form:"pageSize"`
	//OrderBy    string            `json:"order_by"`
	//Conditions map[string]string `form:"conditions"`
	//Searches   map[string]string `form:"searches"`
}

func (service *DeviceListService) List() serializer.Response {
	var res []model.Device
	rtu := make([]*vo.DeviceVo, 0)
	var total int64 = 0
	tx := model.DB.Model(&model.Device{})

	tx.Count(&total)
	tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize).Find(&res)
	for _, d := range res {
		rtu = append(rtu, vo.DeviceDoToVo(&d))
	}

	return serializer.Response{Data: map[string]interface{}{
		"total": total,
		"items": rtu,
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
	err = model.DB.Model(device).Updates(map[string]interface{}{
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
	ConnType string `json:"connType,default=link"`
	RegionID int    `json:"regionId"`
}

func (service DeviceCreateService) Create() serializer.Response {
	var (
		uid uuid.UUID
		err error
	)
	//if model.IsExistsCode(service.Code) {
	//	return serializer.Err(serializer.CodeDeviceCodeRepeatErr, "授权id已存在", nil)
	//}
	device := model.Device{
		ConnType: service.ConnType,
		Name:     service.Name,
		Status:   0,
	}
	if uid, err = uuid.NewUUID(); err != nil {
		return serializer.Err(serializer.CodeRedisError, "生成授权码失败", err)
	}
	device.Code = uid.String()
	if service.ConnType == "gateway" {
		device.RegionId = service.RegionID
	}
	err = device.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建设备失败", err)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"id": device.ID,
		},
	}
}
func (service DeviceCreateService) List() serializer.Response {
	return serializer.Response{}
}

type DeviceDeleteService struct {
	ID uint `json:"id" uri:"id"`
}

// Delete 删除设备
func (service DeviceDeleteService) Delete() serializer.Response {
	device, err := model.GetDeviceByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取设备失败", err)
	}
	if device == nil {
		return serializer.Err(serializer.CodeNotFindDeviceErr, "没有找到指定的设备", err)
	}
	if err := model.DeleteDeviceResource(service.ID); err != nil {
		return serializer.Err(serializer.CodeDBError, "删除设备对应的资源失败", err)
	}
	if err := device.Delete(); err != nil {
		return serializer.Err(serializer.CodeDBError, "删除设备失败", err)
	}
	return serializer.Response{}
}

type DeviceGetService struct {
	ID uint `uri:"id"`
}

func (service DeviceGetService) Get() serializer.Response {
	device, err := model.GetDeviceByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取设备失败", err)
	}
	if device == nil {
		return serializer.Err(serializer.CodeNotFindDeviceErr, "没有找到指定的设备", err)
	}
	return serializer.Response{
		Data: vo.DeviceDoToVo(device),
	}
}
