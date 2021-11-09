package model

import "time"

type Node struct {
	Base
	Name           string    `json:"name" form:"name"`
	Mark           string    `json:"mark" form:"mark"`
	Register       bool      `json:"register" form:"register"`
	LastOnlineTime time.Time `json:"last_online_time" form:"last_online_time"`
}

func (service *Node) TableName() string {
	return "node"
}
