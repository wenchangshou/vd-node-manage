package model

import "gorm.io/gorm"

// DeviceResource 设备资源
type DeviceResource struct {
	gorm.Model
	DeviceID   uint   `json:"device_id"`
	ResourceID string `json:"resource_id"`
}
