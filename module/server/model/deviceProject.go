package model

import "gorm.io/gorm"

// DeviceProject  设备资源
type DeviceProject struct {
	gorm.Model
	DeviceID  uint `json:"device_id"`
	ProjectID uint `json:"_"`
	Project   Project
}

func (dr DeviceProject) TableName() string {
	return "device_project"
}
func (dr DeviceProject) Add() error {
	return DB.Model(&DeviceProject{}).Create(&dr).Error
}
func (dr DeviceProject) Delete() error {
	return DB.Delete(&dr).Error
}

func IsDeviceProject(deviceID uint, projectID uint) bool {
	var count int64
	err := DB.Model(&DeviceProject{}).Where("device_id=? AND project_id=?", deviceID, projectID).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

// GetDeviceProjects 通过设备id和资源id来检索资源
func GetDeviceProjects(deviceID uint, projectIDS []uint) ([]DeviceProject, error) {
	var res []DeviceProject
	err := DB.Model(&DeviceProject{}).Where("device_id=? AND project_id in ?", deviceID, projectIDS).Preload("Project").Find(&res).Error
	return res, err
}

// GetDeviceProject 通过设备id和资源id来检索资源
func GetDeviceProject(deviceID uint, projectId uint) (*DeviceProject, error) {
	var res DeviceProject
	err := DB.Model(&DeviceProject{}).Where("device_id=? AND project_id =? ", deviceID, projectId).Preload("Project").First(&res).Error
	return &res, err
}

func GetDeviceProjectsByDeviceID(deviceID uint) ([]DeviceProject, error) {
	var res []DeviceProject
	err := DB.Model(&DeviceProject{}).Where("device_id=?", deviceID).Preload("Project").Find(&res).Error
	return res, err
}

// DeleteDeviceProject 删除设备资源
func DeleteDeviceProject(deviceID uint) error {
	r := DeviceProject{}
	return DB.Delete(&r, "device_id=?", deviceID).Error
}
func DeleteProjectByDeviceIdAndProjectId(deviceID uint, projectID uint) error {
	return DB.Where("device_id = ? and project_id = ?", deviceID, projectID).Delete(&DeviceProject{}).Error
}
