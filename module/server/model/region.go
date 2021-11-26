package model

import "time"

type Region struct {
	Base
	Name string `gorm:"name" json:"name"`
	LastOnlineTime time.Time `gorm:"last_online_time" json:"last_online_time"`
}
func (Region) TableName()string{
	return "region"
}