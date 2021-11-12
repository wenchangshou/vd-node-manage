package model

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
)

type ProjectRelease struct {
	Base
	Tag       string `json:"number"`
	Content   string `json:"content"`
	Mode      string `json:"mode"`
	Depend    string `json:"depend"`
	Arguments string `json:"arguments"`
	Control   string `json:"control"`
	ProjectID string `json:"_"`
	FileID    string `json:"_"`
	File      File
	Project   Project
}

func (projectRelease *ProjectRelease) Create() (string, error) {
	if err := DB.Create(projectRelease).Error; err != nil {
		logging.GLogger.Warn(fmt.Sprintf("无法插入项目版本：%s", err))
		return "", err
	}
	return projectRelease.ID, nil
}
func (projectRelease *ProjectRelease) Delete() error {
	return DB.Delete(&ProjectRelease{}, projectRelease.ID).Error
}
func GetProjectReleaseByID(id string) (ProjectRelease, error) {
	var projectRelease ProjectRelease
	result := DB.Debug().Model(ProjectRelease{}).Joins("File").Joins("Project").Where("project_releases.ID = ?", id).First(&projectRelease)
	return projectRelease, result.Error
}

func GetProjectReleaseByIdAndProjectId(projectID string, projectReleaseID string) (ProjectRelease, error) {
	var projectRelease ProjectRelease
	result := DB.Debug().Model(ProjectRelease{}).Where("project_id = ? AND project_releases.id = ?", projectID, projectReleaseID).Joins("File").Joins("Project").First(&projectRelease)
	return projectRelease, result.Error

}

// GetProjectReleaseListByProjectID 获取项目所有的发行信息
func GetProjectReleaseListByProjectID(id string) ([]ProjectRelease, error) {
	projectList := make([]ProjectRelease, 0)
	result := DB.Debug().Model(ProjectRelease{}).Where("project_id = ?", id).Joins("File").Find(&projectList)
	return projectList, result.Error
}
