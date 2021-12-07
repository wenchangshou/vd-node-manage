package model

type TaskStatus int
type EventStatus int

const (
	Initializes TaskStatus = iota
	Progress
	Done
	Error
)
const (
	InstallProjectAction EventStatus = iota
	InstallResourceAction
	UpgradeProjectAction
	DeleteResource
	DeleteProject
)

type Event struct {
	ID       uint                   `json:"id"`
	Name     string                 `json:"name"`
	Active   bool                   `json:"active"`
	DeviceID uint                   `json:"deviceID"`
	Action   EventStatus            `json:"action" `
	Status   TaskStatus             `json:"status" `
	Params   map[string]interface{} `json:"params" `
}

type QueryDeviceEventRequest struct {
	DeviceID uint `json:"device_id"`
}
type QueryDeviceEventResponse struct {
	DeviceID uint    `json:"device_id"`
	Count    int     `json:"count"`
	Events   []Event `json:"events"`
}

type DeviceSetEventStatusRequest struct {
	EventID uint        `json:"event_id"`
	Status  EventStatus `json:"status"`
}
