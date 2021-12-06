package model

import (
	"gorm.io/gorm"
)

type ResourceDistribution struct {
	gorm.Model
	DeviceID   uint    `json:"device_id" gorm:"device_id"`
	ResourceID uint    `json:"resource_id" gorm:"resource_id"`
	Schedule   float64 `json:"schedule" gorm:"schedule"`
	Status     int     `json:"status" gorm:"status" gorm:"status"`
	Resource   Resource
}

func (m ResourceDistribution) TableName() string {
	return "resource_distribution"
}

func (m ResourceDistribution) Add() error {
	return DB.Create(&m).Error
}

//func QueryResourceDistributionByDeviceID(id uint) (tasks []dto.Task, err error) {
//var items []ResourceDistribution
//err = DB.Debug().Model(&ResourceDistribution{}).Where("device_id=?", id).Preload("Resource").Find(&items).Error
//if err != nil {
//	return nil, err
//}
//for _, item := range items {
//	params := make(map[string]interface{})
//	params["uri"] = item.Resource.Uri
//	params["status"] = 0
//	params["resource_id"] = item.ResourceID
//	_item := dto.Task{
//		Action: "resourceDistribution",
//		ID:     id,
//		Params: params,
//	}
//	tasks = append(tasks, _item)
//}
//
//fmt.Println(items)
//return
//}

// SetResourceDistributionStatus  设置任务的状态
func SetResourceDistributionStatus(taskId []uint, status uint) error {
	result := DB.Model(&ResourceDistribution{}).Where("id like   ? ", taskId).Update("status", status)
	return result.Error
}
