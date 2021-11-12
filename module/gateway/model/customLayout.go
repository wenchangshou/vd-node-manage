package model

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
)

// 自定义布局
type CustomLayout struct {
	Base
	Name       string `json:"name"`
	Type       string `json:"type"`
	ComputerID string `json:"custom_id"`
	Content    string `json:"content"`
	Sort       int    `json:"sort"`
}

func (customLayout *CustomLayout) TableName() string {
	return "custom_layout"
}
func (customLayout *CustomLayout) Create() (string, error) {
	maxId := GetCustomLayoutMaxId()
	customLayout.Sort = int(maxId) + 1
	if err := DB.Create(customLayout).Error; err != nil {
		logging.GLogger.Warn(fmt.Sprintf("添加新的自定义布局失败:%s", err.Error()))
		return "", err
	}
	return customLayout.ID, nil
}

// GetComputerCustomLayout 获取计算机自定义布局
func (customLayout *CustomLayout) GetComputerCustomLayout() ([]CustomLayout, error) {
	var layouts []CustomLayout
	result := DB.Debug().Model(&CustomLayout{}).Where("computer_id=?", customLayout.ComputerID).Find(&layouts)
	return layouts, result.Error
}
func GetCustomLayoutMaxId() float64 {
	var result float64
	row := DB.Table("custom_layout").Select("max(sort)").Row()
	row.Scan(&result)
	return result
}
