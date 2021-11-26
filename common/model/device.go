package model

// DeviceReportRequest  客户端上报信息
type DeviceReportRequest struct {
	ID       string
	Hostname string
	Ip       string
}
type DeviceRegisterRequest struct {
	Code string `json:"code"`
	ConnType string `json:"connType"`
}
type DeviceRegisterResponse struct {
	ID     uint `json:"id"`
	Status int    `json:"status"`
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type DeviceUpdateInfo struct {
	LastUpdate    int64
	ReportRequest *DeviceReportRequest
}

type DeviceStatusType int32

const (
	Device_Init = iota
	Device_Register
	Device_Disable
)
