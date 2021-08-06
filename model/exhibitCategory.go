package model

import "gorm.io/gorm"

type ExhibitCategory struct {
	gorm.Model
	Name string
}
