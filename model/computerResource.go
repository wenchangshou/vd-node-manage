package model

import "gorm.io/gorm"

type ComputerResource struct {
	gorm.Model
	ComputerID uint `gorm:"computer_id"`
	ResourceID uint `gorm:"resource_id"`
	Resource   Resource
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
func GetComputerResourceById(id int) (*ComputerResource, error) {
	computerResource := &ComputerResource{}
	result := DB.First(computerResource, id)
	return computerResource, result.Error
}

// GetComputerResourceByComputerId 通过计算机id来获取指定计算机资源
func GetComputerResourceByComputerId(id int) ([]ComputerResource, error) {
	var computerResourceList []ComputerResource
	result := DB.Where("computer_id = ?", id).Joins("Resource").Find(&computerResourceList)
	return computerResourceList, result.Error
}
func DeleteComputerResourceById(id int) error {
	result := DB.Debug().Delete(&ComputerResource{}, id)
	return result.Error
}
