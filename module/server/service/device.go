package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/wenchangshou/vd-node-manage/common/cache"
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
	"github.com/wenchangshou/vd-node-manage/module/server/vo"
	"strconv"
	"strings"
)

type DeviceListService struct {
	Page       int               `json:"page" binding:"min=1,required" form:"page"`
	PageSize   int               `json:"pageSize" binding:"min=1,required" form:"pageSize"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}

func (service *DeviceListService) List() serializer.Response {
	var res []model.Device
	var ids []string
	rtu := make([]*vo.DeviceVo, 0)
	var total int64 = 0
	span := opentracing.StartSpan("list device")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	session := model.DB.WithContext(ctx)
	tx := session.Model(&model.Device{})
	for k, v := range service.Conditions {
		if k == "deviceId" {
			k = "id"
		}
		arr := strings.Split(v, ",")
		if len(arr) > 1 {
			tx = tx.Where(k+" IN ?", arr)
			continue
		}
		tx = tx.Where(k+" = ?", v)
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
	for _, d := range res {
		i := strconv.Itoa(int(d.ID))
		ids = append(ids, i)
	}
	m, _ := cache.GetSettings(ids, "device-")
	for _, d := range res {
		o := vo.DeviceDoToVo(&d)
		_id := strconv.Itoa(int(d.ID))
		e, exists := m[_id]
		o.Online = false
		if exists {
			m := make(map[string]interface{})
			json.Unmarshal([]byte(e), &m)
			o.LastOnlineTime = int64(m["last_online_time"].(float64))
			o.Detailed = m["body"].(string)
			o.Online = true
		}
		rtu = append(rtu, o)
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
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	ConnType string `json:"connType,default=link"`
	RegionID int    `json:"regionId"`
	Expired  uint64 `json:"expired"`
	Mode     int    `json:"mode"`
}

func (service DeviceCreateService) Create() serializer.Response {
	var (
		uid uuid.UUID
		err error
	)
	device := model.Device{
		ConnType: service.ConnType,
		Name:     service.Name,
		Status:   0,
		Expired:  service.Expired,
		Mode:     service.Mode,
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
	c, exists := cache.Get(fmt.Sprintf("device-%d", service.ID))
	rtu := vo.DeviceDoToVo(device)
	rtu.Online = false
	if exists {
		rtu.Online = true
		m := make(map[string]interface{})
		json.Unmarshal([]byte(c.(string)), &m)
		rtu.LastOnlineTime = int64(m["last_online_time"].(float64))
		rtu.Detailed = m["body"].(string)
	}
	return serializer.Response{
		Data: rtu,
	}
}

type DeviceGetOnlineService struct {
	IDS []uint `json:"id" uri:"id" form:"id"`
}
type DeviceGetOnlineRtu struct {
	ID     uint `json:"id"`
	Online bool `json:"online"`
}

func (service DeviceGetOnlineService) Get() serializer.Response {
	arr := make([]string, 0)
	for _, v := range service.IDS {
		tmp := strconv.Itoa(int(v))
		arr = append(arr, tmp)
	}
	rtu := make([]DeviceGetOnlineRtu, 0)
	for ids := range arr {
		id, _ := strconv.Atoi(arr[ids])
		r := DeviceGetOnlineRtu{Online: false, ID: uint(id)}
		rtu = append(rtu, r)
	}
	k, _ := cache.GetSettings(arr, "device-")
	for k2 := range rtu {
		_id := strconv.Itoa(int(rtu[k2].ID))
		if _, exists := k[_id]; exists {
			rtu[k2].Online = true
		}
	}
	return serializer.Response{
		Data: rtu,
	}
}

type UpdateDeviceStruct struct {
	ID      []uint `json:"id" uri:"id" form:"id"`
	Mode    int    `json:"mode" form:"mode"`
	Expired uint64 `json:"expired"`
}

func (service *UpdateDeviceStruct) SetLease() serializer.Response {
	var (
		err error
	)
	maps := make(map[string]interface{})
	if service.Mode == model2.LEASE_ENABLE && service.Expired == 0 {
		return serializer.Err(serializer.CodeParamErr, "enable lease,expired不能為空", nil)
	}
	maps["expired"] = service.Expired
	maps["mode"] = service.Mode
	err = model.SetDevices(service.ID, maps)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "'设置设备失败", err)
	}
	return serializer.Response{}
}
