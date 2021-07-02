package model

import "gorm.io/gorm"

type ComputerProject struct {
	gorm.Model
	ComputerId       uint `gorm:"computer_id"`
	ProjectId        uint `gorm:"project_id"`
	ProjectReleaseId uint `gorm:"project_release_id"`
	Status           uint `gorm:"status"`
}

func GetComputerProjectById(cid uint, pid uint) (ComputerProject, error) {
	var computerProject ComputerProject
	result := DB.Debug().Model(&ComputerProject{}).Where("computer_id=? AND project_id  = ?", cid, pid).First(&computerProject)
	return computerProject, result.Error
}
