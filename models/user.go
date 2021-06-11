package models

import "gorm.io/gorm"

const (
	// Active 账户正常状态
	Active = iota
	// NotActivicated 未激活
	NotActivicated
	// Baned 被封禁
	Baned
	// OveruseBaned 超额使用被封禁
	OveruseBaned
)

type User struct {
	// 表字段
	gorm.Model
	Username string `gorm:"size:50"`
	Password string `json:"-"`
	Status   int
	Type     int
}
