package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ID        int
}

func (base *Base) BeforeCreate(_ *gorm.DB) (err error) {
	return
}
