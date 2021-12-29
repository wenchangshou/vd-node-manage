package model

type Window struct {
	ID        string `json:"id"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Z         int    `json:"z"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Service   string `json:"service"`
	Style     map[string]interface{}
	Arguments map[string]interface{}
	Source    struct {
		Category string `json:"category"`
		Path     string `json:"path"`
		URI      string `json:"uri"`
		ID       uint   `json:"id"`
	}
}
type OpenWindowInfo struct {
	ID        string `json:"id"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Z         int    `json:"z"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Service   string `json:"service"`
	Style     map[string]interface{}
	Arguments map[string]interface{}
	Source    string `json:"source"`
}
type ActiveWindowInfo struct {
	ID string `json:"id"`
	//Code int    `json:"code"`
	//Msg  string `json:"msg"`
	Info string `json:"info"`
	Run  bool   `json:"run"`
}
