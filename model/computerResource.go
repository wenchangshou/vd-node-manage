package model

import "gorm.io/gorm"

type ComputerResource struct {
	gorm.Model
	ComputerId uint `gorm:"computer_id"`
	ResourceId uint `gorm:"resource_id"`
	Status     uint `gorm:"status"`
}

func (ComputerResource) TableName() string {
	return "computer_resource"
}
func (computerResource *ComputerResource) Create() (uint, error) {
	if err := DB.Create(computerResource).Error; err != nil {
		return 0, err
	}
	return computerResource.ID, nil
}
