package model

import (
	"time"

	"gorm.io/gorm"
)

type Computer struct {
	gorm.Model
	Name           string    `gorm:"name"`
	Ip             string    `gorm:"ip"`
	Mac            string    `gorm:"mac" validate:"required,mac"`
	HostName       string    `gorm:"hostname"`
	LastOnlineTime time.Time `gorm:"last_online_time"`
}

func (Computer) TableName() string {
	return "computer"
}
func AddClient(data map[string]interface{}) error {
	return nil
}
func (computer *Computer) IsExistByMac() bool {
	var client2 Computer
	err := DB.Debug().Select("mac").Where("mac = ?", computer.Mac).First(&client2).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (computer *Computer) UpdateByMac() error {
	data := make(map[string]interface{})
	data["ip"] = computer.Ip
	data["mac"] = computer.Mac
	data["host_name"] = computer.HostName
	data["last_online_time"] = time.Now()
	return DB.Model(&Computer{}).Where("mac=?", computer.Mac).Updates(data).Error
}

func (computer *Computer) Create() error {
	_client := Computer{
		Ip:       computer.Ip,
		HostName: computer.HostName,
		Mac:      computer.Mac,
	}
	return DB.Create(&_client).Error
}

// GetComputerByMac 通过mac地址获取用户信息
func GetComputerByMac(mac string) (Computer, error) {
	var computer Computer
	result := DB.Model(&Computer{}).Where("mac = ?", mac).First(&computer)
	return computer, result.Error
}
