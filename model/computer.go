package model

import (
	"time"

	"gorm.io/gorm"
)

type Computer struct {
	gorm.Model
	Source         string    `gorm:"source" json:"source"`
	Switchs        string    `gorm:"switchs" json:"switchs"`
	Active         bool      `gorm:"active" json:"active"`
	Open           bool      `gorm:"open" json:"open"`
	MenuIndex      int       `gorm:"menu_index" json:"menu_index"`
	LayoutIndex    int       `gorm:"layout_index" json:"layout_index"`
	SelectedNum    string    `gorm:"selected_num" json:"selected_num" default:"[]"`
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
func UpdateComputerById(id int, data map[string]interface{}) error {
	return DB.Model(&Computer{}).Where("id = ?", id).Updates(data).Error
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
func GetComputerById(id int) (Computer, error) {
	var computer Computer
	result := DB.Model(&Computer{}).First(&computer, id)
	return computer, result.Error

}
func ListComputer() ([]Computer, int64) {
	var (
		computers []Computer
		total     int64
	)
	DB.Model(&Computer{}).Count(&total)
	DB.Find(&computers)
	return computers, total
}

func GetComputerDetailsById(id int) {

}
