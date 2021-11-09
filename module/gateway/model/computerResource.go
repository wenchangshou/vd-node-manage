package model

// type ComputerResource struct {
// 	Base
// 	ComputerID string     `gorm:"computer_id" gorm:"primaryKey"`
// 	ResourceID string     `gorm:"resource_id" gorm:"primaryKey"`
// 	Resource   []Resource `gorm:"many2many:Computecomputer_resource"`
// 	Status     uint       `gorm:"status"`
// }

// func (ComputerResource) TableName() string {
// 	return "computer_resource"
// }
// func (computerResource *ComputerResource) Create() (string, error) {
// 	if err := DB.Create(computerResource).Error; err != nil {
// 		return "", err
// 	}
// 	return computerResource.ID, nil
// }
// func GetComputerResourceById(id int) (*ComputerResource, error) {
// 	computerResource := &ComputerResource{}
// 	result := DB.First(computerResource, id)
// 	return computerResource, result.Error
// }
// func GetComputerResourceByComputerIdAndResourceId(computerId string, resourceId string) (*ComputerResource, error) {
// 	computerResource := &ComputerResource{}
// 	result := DB.Where("computer_id = ? AND resource_id = ?", computerId, resourceId).First(&computerResource)
// 	return computerResource, result.Error

// }

// // GetComputerResourceByComputerId 通过计算机id来获取指定计算机资源
// func GetComputerResourceByComputerId(id string) ([]ComputerResource, error) {
// 	var computerResourceList []ComputerResource
// 	result := DB.Where("computer_id = ?", id).Joins("Resource").Find(&computerResourceList)
// 	return computerResourceList, result.Error
// }
// func DeleteComputerResourceById(id int) error {
// 	result := DB.Debug().Delete(&ComputerResource{}, id)
// 	return result.Error
// }

// // 通过项目id列表来批量获取
// func GetComputerResourcesByIds(ids []string) ([]ComputerResource, error) {
// 	var computerResources []ComputerResource
// 	result := DB.Debug().Where("resource_id IN ?", ids).Find(&computerResources)
// 	return computerResources, result.Error
// }

// // 通过项目id列表来批量获取
// func GetComputerResourcesByComputerIds(ids []string) ([]ComputerResource, error) {
// 	var computerResources []ComputerResource
// 	result := DB.Debug().Where("computer_id IN ?", ids).Find(&computerResources)
// 	return computerResources, result.Error
// }
// func GetComputerResource(page int, size int, orderBy string, conditions map[string]string) ([]ComputerResource, int64) {
// 	var res []ComputerResource
// 	var total int64
// 	tx := DB.Model(&ComputerResource{})
// 	if orderBy != "" {
// 		tx = tx.Order(orderBy)
// 	}
// 	for k, v := range conditions {
// 		tx = tx.Where(k+" = ?", v)
// 	}
// 	tx.Count(&total)
// 	tx.Debug().Limit(size).Offset((page - 1) * size).Preload("Resources").Find(&res)
// 	return res, total
// }

// func GetComputerResourceByComputerIDAndResourceCategory(page int, size int, computerID string, category string) ([]Resource, int64) {
// 	var res []Resource
// 	var total int64
// 	tx := DB.Model(&Resource{})
// 	tx.Joins("")
// 	return res, total

// }
