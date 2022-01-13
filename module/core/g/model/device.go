package model

type ServerInfo struct {
	Name     string `json:"name"`
	Register bool   `json:"register"`
	Address  string `json:"address"`
	Expired  int64  `json:"expired"`
	Detailed struct {
		Communication bool `json:"communication"`
		Server        bool `json:"server"`
	} `json:"detailed"`
}
