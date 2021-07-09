package model

import (
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name       string `gorm:"name"`
	Mode       string `gorm:"mode"`
	SourceName string `gorm:"source_name"`
	UserId     uint   `gorm:"user_id"`
	Size       uint   `gorm:"size"`
	Uuid       string `gorm:"uuid"`
}

func (file *File) TableName() string {
	return "file"
}
func (file *File) Create() (uint, error) {
	if err := DB.Create(file).Error; err != nil {
		logging.G_Logger.Warn(fmt.Sprintf("无法插入文件:%s", err))
	}
	return file.ID, nil
}
func GetFileByUidAndId(id int, uid uint) (File, error) {
	var file File
	result := DB.Where("id = ? AND user_id = ?", id, uid).First(&file)
	return file, result.Error
}

func GetFileById(id uint) (File, error) {
	var file File
	result := DB.Where("id = ? ", id).First(&file)
	return file, result.Error

}
