package model

import (
	"fmt"
	"gorm.io/gorm"
)

// Device 设备
type Device struct {
	gorm.Model
	Code           string `gorm:"code" json:"code"`
	ConnType       string `gorm:"conn_type" json:"connType"`
	HardwareCode   string `gorm:"hardware_code" json:"hardware_code"`
	Name           string `gorm:"name" json:"name"`
	HostName       string `gorm:"hostname" json:"hostName"`
	Status         int    `gorm:"status" json:"status"`
	LastOnlineTime int64  `gorm:"last_online_time" json:"last_online_time"`
	RegionId       int    `gorm:"region_id" json:"region_id"`
}

func (Device) TableName() string {
	return "device"
}
func IsExistsCode(code string) bool {
	device := &Device{}
	DB.First(&device, "code=?", code)
	return device.ID > 0
}

// Create 创建一个新的客户端
func (device *Device) Create() error {
	err := DB.Debug().Model(&Device{}).Create(device).Error
	fmt.Println(err)
	return err
}
func GetDeviceByID(id uint) (*Device, error) {
	var device *Device
	result := DB.Where("id=?", id).First(&device)
	return device, result.Error
}
