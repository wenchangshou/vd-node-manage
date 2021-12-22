package model

// DeviceSetStatusRequest 设置任务状态
type DeviceSetStatusRequest struct {
	ID     []uint      `json:"id"`
	Status EventStatus `json:"status"`
}

// DeviceReportRequest  客户端上报信息
type DeviceReportRequest struct {
	ID   uint
	Info string
	//Hostname string
	//Ip       string
}
type DeviceRegisterRequest struct {
	Code         string `json:"code"`
	ConnType     string `json:"connType"`
	HardwareCode string `json:"hardwareCode"`
}
type DeviceRegisterResponse struct {
	ID           uint   `json:"id"`
	Status       int    `json:"status"`
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	RpcAddress   string `json:"rpcAddress"`
	HttpAddress  string `json:"httpAddress"`
	RedisAddress string `json:"redis_address"`
}

// DeviceQueryStatusRequest 设备查询任务请求结构体
type DeviceQueryStatusRequest struct {
	ID uint `json:"id"`
}

// DeviceQueryStatusResponse 设备查询任务返回结构体
type DeviceQueryStatusResponse struct {
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

type DeviceAddResourceRequest struct {
	ID         uint `json:"id"`
	ResourceID uint `json:"resource_id"`
}
type DeviceDeleteResourceRequest struct {
	ID         uint `json:"id"`
	ResourceID uint `json:"resource_id"`
}
