package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ID        string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" `
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID != "" {
		return nil

	}
	uuid2, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.ID = uuid2.String()
	return
}
