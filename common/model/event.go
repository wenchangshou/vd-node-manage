package model

import "time"

type Event struct {
	Action    string    `json:"active"`
	ID        string    `json:"id"`
	Arguments string    `json:"arguments"`
	T         time.Time `json:"t"`
}
