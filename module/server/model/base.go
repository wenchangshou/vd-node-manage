package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ID        int
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	//var (
	//	uuid2 uuid.UUID
	//)
	//if base.ID != "" {
	//	return nil
	//
	//}
	//uuid2, err = uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	//base.ID = uuid2.String()
	return
}