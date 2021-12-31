package model

import (
	"gorm.io/gorm"
	"strings"
)

type Resource struct {
	gorm.Model
	Name     string `json:"name" gorm:"name"`
	Service  string `json:"service"`
	Category string `gorm:"category" json:"category"`
	Uri      string `json:"uri"`
	Status   int    `json:"status"`
}

func (resources *Resource) TableName() string {
	return "resource"
}
func (resources *Resource) Delete() error {
	return DB.Model(&Resource{}).Where("id=?", resources.ID).Delete(resources).Error
}
func (resources *Resource) Create() (uint, error) {
	if err := DB.Create(resources).Error; err != nil {
		return 0, err
	}
	return resources.ID, nil
}

func (resources *Resource) Add() (uint, error) {
	err := DB.Create(&resources).Error
	return resources.ID, err
}
func GetResourceById(id uint) (*Resource, error) {
	var resource *Resource
	result := DB.Model(&Resource{}).Where("id=?", id).First(&resource)
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
			search += k + " like '%" + v + "%' OR "
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	tx.Debug().Limit(size).Offset((Page - 1) * size).Joins("File").Find(&res)
	return res, total
}

func GetResourcesByIds(ids []uint) ([]Resource, error) {

	var res []Resource

	// err := DB.Model(&Computer{}).Association("Resources").DB.Where("computer_id=?", ids).Find(&res).Error
	err := DB.Model(&Resource{}).Where("id like ?", ids).Find(&res).Error
	return res, err
}
func ListResource(Page int, size int, orderBy string, conditions map[string]string, searches map[string]string) (resources []Resource, total int64) {
	tx := DB.Model(Resource{})
	if orderBy != "" {
		tx = tx.Order(orderBy)
	}
	for k, v := range conditions {
		tx = tx.Where(k+" = ?", v)
	}
	if len(searches) > 0 {
		search := ""
		for k, v := range searches {
			search += k + " like '%" + v + "%' OR "
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	tx.Limit(size).Offset((Page - 1) * size).Preload("Computers").Find(&resources)
	return
}
