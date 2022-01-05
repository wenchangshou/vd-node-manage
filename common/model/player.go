package model

type Player struct {
	Startup    string `json:"startup"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Version    string `json:"version"`
	UpdateTime int64  `json:"update_time"`
}
