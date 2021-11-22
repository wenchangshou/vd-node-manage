package model

// DeviceReportRequest  客户端上报信息
type DeviceReportRequest struct {
	ID       string
	Hostname string
	Ip       string
}
type DeviceRegisterRequest struct {
	HID string `json:"hid"`
	Type string `json:"type"`
	Params string `json:"params"`
}
type DeviceRegisterResponse struct {
	ID string `json:"id"`
	Status int `json:"status"`
}