package model

import (
	"strings"

	"gorm.io/gorm"
)

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
func (resources *Resource) Delete() error {
	return DB.Delete(&Resource{}, resources.ID).Error
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

func GetResources(Page int, size int, orderBy string, conditions map[string]string, searches map[string]string) ([]Resource, int64) {
	var res []Resource
	var total int64
	tx := DB.Model(&Resource{})
	if orderBy != "" {
		tx = tx.Order(orderBy)
	}
	for k, v := range conditions {
		tx = tx.Where(k+" = ?", v)
	}
	if len(searches) > 0 {
		search := ""
		for k, v := range searches {
			search += (k + " like '%" + v + "%' OR ")
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	tx.Debug().Limit(size).Offset((Page - 1) * size).Joins("File").Find(&res)
	return res, total
}
