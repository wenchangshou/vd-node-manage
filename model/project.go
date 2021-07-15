package model

import (
	"fmt"
	"strings"

	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Start       string `gorm:"start"`
	Name        string `gorm:"size:50"`
	Category    string `gorm:"category"`
	Description string `gorm:"description"`
	Arguments   string `gorm:"arguments"`
	Control     string `gorm:"control"`
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
