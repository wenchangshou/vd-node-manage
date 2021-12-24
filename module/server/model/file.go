package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name       string `gorm:"name" json:"name"`
	Category   string
	Mode       string `gorm:"mode" json:"mode"`
	SourceName string `gorm:"source_name" json:"sourceName"`
	Size       uint   `gorm:"size" json:"size"`
	Uuid       string `gorm:"uuid" json:"uuid"`
	Md5        string `gorm:"md5" json:"md5"`
}

func (file *File) TableName() string {
	return "file"
}
func (file *File) Create() (uint, error) {
	if err := DB.Create(file).Error; err != nil {
		return 0, err
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
