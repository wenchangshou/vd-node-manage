package models

import "gorm.io/gorm"

type ProjectRelease struct {
	gorm.Model
	Number  string `json:"number"`
	Content string `json:"content"`
	Mode    string `json:"mode"`
	Depend  uint   `json:"depend"`
}
