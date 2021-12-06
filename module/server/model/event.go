package model

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name     string            `json:"name" gorm:"name"`
	Active   bool              `gorm:"active" json:"active"`
	DeviceID uint              `json:"deviceID" gorm:"deviceID"`
	Action   model.EventStatus `json:"action" gorm:"action"`
	Status   model.TaskStatus  `json:"status" gorm:"status"`
	Params   string            `json:"params" gorm:"params"`
}

func (e Event) TableName() string {
	return "event"
}
func (e Event) Add() error {
	return DB.Create(&e).Error
}
func QueryDeviceEventByDeviceID(id uint) (tasks []Event, err error) {
	var items []Event
	err = DB.Debug().Model(&Event{}).Where("device_id=?", id).Find(&items).Error
	return items, err
}
