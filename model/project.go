package model

import (
	"fmt"
	"strings"

	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Start         string `gorm:"start" json:"start"`
	Name          string `gorm:"size:50" json:"name"`
	Category      string `gorm:"category" json:"category"`
	Description   string `gorm:"description" json:"description"`
	Arguments     string `gorm:"arguments" json:"arguments"`
	Control       string `gorm:"control" json:"control"`
	Cover         int    `gorm:"cover" json:"_"`
	CoverImageUrl string `json:"cover_image_url"`
	File          File   `gorm:"foreignKey:Cover" json:"_"`
}

func (project *Project) TableName() string {
	return "project"
}

// Create 创建一个项目
func (project *Project) Create() (uint, error) {
	if err := DB.Create(project).Error; err != nil {
		logging.G_Logger.Warn(fmt.Sprintf("无法插入离线下载任务：%s", err))
		return 0, err
	}
	return project.ID, nil
}

//DeleteProjectByID 删除项目
func DeleteProjectByID(id, uid uint) error {
	result := DB.Where("id = ? and user_id=?", id, uid).Delete(&Project{})
	return result.Error
}
func GetProjects(Page int, size int, orderBy string, conditions map[string]string, searches map[string]string) ([]Project, int64) {
	var res []Project
	var total int64
	tx := DB.Model(&Project{})
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
	tx.Debug().Limit(size).Offset((Page - 1) * size).Find(&res)
	return res, total
}
func GetProjectByIds(ids []int) ([]Project, error) {
	var projects []Project
	result := DB.Debug().Model(&Project{}).Where("project.id in  ? ", ids).Joins("File").Find(&projects)
	for key, project := range projects {
		if project.File.ID > 0 {
			projects[key].CoverImageUrl = "upload/" + project.File.SourceName
		}
	}
	return projects, result.Error

}
