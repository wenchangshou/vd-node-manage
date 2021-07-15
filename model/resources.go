package model

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name     string `gorm:"name"`
	Category string `gorm:"category"`
	FileID   uint   `json:"_"`
	File     File
}

func (resources *Resource) TableName() string {
	return "resources"
}
func (resources *Resource) Create() (int, error) {
	if err := DB.Create(resources).Error; err != nil {
		return -1, err
	}
	return int(resources.ID), nil
}
func GetResourceById(id uint) (*Resource, error) {
	var resource *Resource
	result := DB.Debug().Model(&Resource{}).Joins("File").Where("resources.id=?", id).First(&resource)
	return resource, result.Error
}
