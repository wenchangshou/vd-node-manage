package model

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
)

type File struct {
	Base
	Name       string `gorm:"name" json:"name"`
	Mode       string `gorm:"mode" json:"mode"`
	SourceName string `gorm:"source_name" json:"sourceName"`
	UserId     string `gorm:"user_id" json:"user_id"`
	Size       uint   `gorm:"size" json:"size"`
	Uuid       string `gorm:"uuid" json:"uuid"`
	Md5        string `gorm:"md5" json:"md5"`
}

func (file *File) TableName() string {
	return "file"
}
func (file *File) Create() (string, error) {
	if err := DB.Create(file).Error; err != nil {
		logging.GLogger.Warn(fmt.Sprintf("无法插入文件:%s", err))
	}
	return file.ID, nil
}
func (file *File) Delete() error {
	return DB.Where("id=?", file.ID).Delete(&file).Error
}
func GetFileByUidAndId(id string, uid string) (File, error) {
	var file File
	result := DB.Where("id = ? AND user_id = ?", id, uid).First(&file)
	return file, result.Error
}

func GetFileById(id string) (File, error) {
	var file File
	result := DB.Where("id = ? ", id).First(&file)
	return file, result.Error
}
