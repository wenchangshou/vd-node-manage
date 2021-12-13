package model

import "gorm.io/gorm"

// DeviceResource 设备资源
type DeviceResource struct {
	gorm.Model
	DeviceID   uint `json:"device_id"`
	ResourceID uint `json:"resource_id"`
}

func (dr DeviceResource) TableName() string {
	return "device_resource"
}
func (dr DeviceResource) Add() error {
	return DB.Model(&DeviceResource{}).Create(&dr).Error
}
func (dr DeviceResource) Delete() error {
	return DB.Delete(&dr).Error
}

func IsDeviceResource(deviceID uint, resourceID uint) bool {
	var count int64
	err := DB.Model(&DeviceResource{}).Where("device_id=? AND resource_id=?", deviceID, resourceID).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
