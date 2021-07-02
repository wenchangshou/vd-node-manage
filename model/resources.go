package model

import "gorm.io/gorm"

type Resources struct {
	gorm.Model
	Name     string `gorm:"name"`
	Category string `gorm:"category"`
	FileId   int    `gorm:"file_id"`
}

func (resources *Resources) TableName() string {
	return "resources"
}
func (resources *Resources) Create() (int, error) {
	if err := DB.Create(resources).Error; err != nil {
		return -1, err
	}
	return int(resources.ID), nil
}
