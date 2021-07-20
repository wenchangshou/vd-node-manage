package model

import (
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"gorm.io/gorm"
)

type ProjectRelease struct {
	gorm.Model
	Tag       string `json:"number"`
	Content   string `json:"content"`
	Mode      string `json:"mode"`
	Depend    uint   `json:"depend"`
	Arguments string `json:"arguments"`
	Control   string `json:"control"`
	ProjectID uint   `json:"_"`
	FileID    uint   `json:"_"`
	File      File
	Project   Project
}

func (projectRelease *ProjectRelease) Create() (uint, error) {
	if err := DB.Create(projectRelease).Error; err != nil {
		logging.G_Logger.Warn(fmt.Sprintf("无法插入项目版本：%s", err))
		return 0, err
	}
	return projectRelease.ID, nil
}
func GetProjectReleaseByID(id uint) (ProjectRelease, error) {
	var projectRelease ProjectRelease
	result := DB.Debug().Model(ProjectRelease{}).Joins("File").Joins("Project").First(&projectRelease)
	return projectRelease, result.Error
}

func GetProjectReleaseByIdAndProjectId(projectID uint, projectReleaseID uint) (ProjectRelease, error) {
	var projectRelase ProjectRelease
	result := DB.Debug().Model(ProjectRelease{}).Joins("File").Joins("Project").First(&projectRelase)
	return projectRelase, result.Error

}
func GetProjectReleaseListByProjectID(id uint) ([]ProjectRelease, error) {
	projectList := make([]ProjectRelease, 0)
	result := DB.Debug().Model(ProjectRelease{}).Where("project_id = ?", id).Joins("File").Find(&projectList)
	return projectList, result.Error
}
