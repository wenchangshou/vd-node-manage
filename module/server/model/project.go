package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name    string `json:"name" gorm:"name"`
	Service string `json:"service"`
	Uri     string `json:"uri"`
	Status  int    `json:"status"`
	Md5     string `json:"md5"`
	Startup string `json:"startup"`
}

func (project Project) TableName() string {
	return "project"
}
func (project *Project) Create() (uint, error) {
	if err := DB.Create(project).Error; err != nil {
		return 0, err
	}
	return project.ID, nil
}
func GetProjectByID(id uint) (project *Project, err error) {
	result := DB.Model(&Project{}).Where("id=?", id).First(&project)
	return project, result.Error
}
