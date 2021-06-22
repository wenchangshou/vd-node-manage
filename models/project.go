package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `gorm:"size:50"`
	Category    string `gorm:"category"`
	Description string `gorm:"description"`
	Arguments   string `gorm:"arguments"`
}
