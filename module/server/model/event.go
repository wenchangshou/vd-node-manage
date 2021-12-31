package model

import (
	"github.com/wenchangshou/vd-node-manage/common/model"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name       string            `json:"name" gorm:"name"`
	Active     bool              `gorm:"active" json:"active"`
	DeviceID   uint              `json:"deviceID" gorm:"deviceID"`
	Action     model.EventAction `json:"action" gorm:"action"`
	Status     model.EventStatus `json:"status" gorm:"status"`
	ResourceId uint              `json:"resourceId" gorm:"resource"`
	ProjectId  uint              `json:"projectID" gorm:"project_id"`
	//Params   string            `json:"params" gorm:"params"`
}

func (e Event) TableName() string {
	return "event"
}
func (e Event) Add() error {
	return DB.Create(&e).Error
}
func QueryDeviceEventByDeviceIDAndStatus(id uint, status model.EventStatus) (tasks []Event, err error) {
	var items []Event
	err = DB.Model(&Event{}).Where("device_id=? AND status =?", id, status).Find(&items).Error
	return items, err
}

func SetDeviceEventStatus(id uint, status model.EventStatus) error {
	return DB.Model(&Event{}).Where("id=?", id).Update("status", status).Error
}
