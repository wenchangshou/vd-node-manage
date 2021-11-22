package model

import "time"

// Device 设备
type Device struct {
	Base
	Type           string    `gorm:"type" json:"type"`
	Name           string    `gorm:"name" json:"name"`
	HostName       string    `gorm:"hostname" json:"hostName"`
	Status         int       `gorm:"status" json:"status"`
	LastOnlineTime time.Time `gorm:"last_online_time" json:"last_online_time"`
	Resource []Resource `gorm:"many2many:computer_resources;"json:"_"`
}

func (Device) TableName() string {
	return "computer"
}
