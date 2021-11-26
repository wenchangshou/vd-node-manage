package model

import (
	"gorm.io/gorm"
	"strings"
)

type Resource struct {
	gorm.Model
	Name string `json:"name" gorm:"name"`
	Category string `gorm:"category" json:"category"`
	Uri string
	FileID uint `json:"_"`
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
func GetResourceById(id string) (*Resource, error) {
	var resource *Resource
	result := DB.Debug().Model(&Resource{}).Joins("File").Where("resources.id=?", id).Preload("Computers").First(&resource)
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

func GetResourcesByIds(ids []string) ([]Resource, error) {

	var res []Resource

	// err := DB.Model(&Computer{}).Association("Resources").DB.Where("computer_id=?", ids).Find(&res).Error
	return res, nil
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
			search += (k + " like '%" + v + "%' OR ")
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	tx.Debug().Limit(size).Offset((Page - 1) * size).Preload("Computers").Find(&resources)
	return
}
