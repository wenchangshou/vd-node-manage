package model

import "encoding/json"

type Player struct {
	Startup    string `json:"startup"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Version    string `json:"version"`
	UpdateTime int64  `json:"update_time"`
}

func (player Player) Serialization() []byte {
	b, _ := json.Marshal(player)
	return b
}
func GeneratePlayer(startup string, name string, path string, version string, updateTime int64) *Player {
	return &Player{
		Startup:    startup,
		Name:       name,
		Path:       path,
		Version:    version,
		UpdateTime: updateTime,
	}
}
