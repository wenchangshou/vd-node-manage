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
