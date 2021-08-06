package model

import "gorm.io/gorm"

type Exhibiton struct {
	gorm.Model
	ExhibitcategoryID int    `gorm:"exhibition_category_id"`
	Name              string `gorm:"name" json:"name"`
	Level             int    `gorm:"level" json:"level"`
	Encryption        string `gorm:"encryption" json:"encryption"`
}
