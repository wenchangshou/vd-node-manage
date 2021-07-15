package model

import (
	"errors"

	"gorm.io/gorm"
)

type ComputerProject struct {
	gorm.Model
	ComputerId       uint `gorm:"computer_id"`
	ProjectId        uint `gorm:"project_id"`
	ProjectReleaseId uint `gorm:"project_release_id"`
	Status           uint `gorm:"status"`
}

func (cp *ComputerProject) Create() (uint, error) {
	if err := DB.Create(cp).Error; err != nil {
		return 0, err
	}
	if cp.ID == 0 {
		return 0, errors.New("未找到相应的记录")
	}
	return cp.ID, nil
}
func GetComputerProjectById(cid uint, pid uint) (ComputerProject, error) {
	var computerProject ComputerProject
	result := DB.Debug().Model(&ComputerProject{}).Where("computer_id=? AND project_id  = ?", cid, pid).First(&computerProject)
	return computerProject, result.Error
}
func DeleteComputerProjectById(id uint32) error {
	result := DB.Debug().Delete(&ComputerProject{}, id)
	return result.Error
}
