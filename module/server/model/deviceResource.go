package model

import "gorm.io/gorm"

// DeviceResource 设备资源
type DeviceResource struct {
	gorm.Model
	DeviceID   uint     `json:"device_id"`
	ResourceID uint     `json:"_"`
	Resource   Resource `gorm:"ForeignKey:ResourceID;AssociationForeignKey:ID"`
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

// GetDeviceResources 通过设备id和资源id来检索资源
func GetDeviceResources(deviceID uint, resourceIDS []uint) ([]DeviceResource, error) {
	var res []DeviceResource
	err := DB.Model(&DeviceResource{}).Where("device_id=? AND resource_id in ?", deviceID, resourceIDS).Preload("Resource").Find(&res).Error
	return res, err
}

func GetDeviceResourcesByDeviceID(deviceID uint) ([]DeviceResource, error) {
	var res []DeviceResource
	err := DB.Model(&DeviceResource{}).Where("device_id=?", deviceID).Preload("Resource").Find(&res).Error
	return res, err
}

// DeleteDeviceResource 删除设备资源
func DeleteDeviceResource(deviceID uint) error {
	r := DeviceResource{}
	return DB.Delete(&r, "device_id=?", deviceID).Error
}
