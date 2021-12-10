package model

type EventStatus int
type EventAction int

const (
	Initializes EventStatus = iota
	Progress
	Done
	Error
	UNKNOWN
	WAITING
)
const (
	InstallProjectAction EventAction = iota
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
	Action   EventAction            `json:"action" `
	Status   EventStatus            `json:"status" `
	Params   map[string]interface{} `json:"params" `
}

type QueryDeviceEventRequest struct {
	DeviceID uint        `json:"device_id"`
	Status   EventStatus `json:"status"`
}
type QueryDeviceEventResponse struct {
	DeviceID uint    `json:"device_id"`
	Count    int     `json:"count"`
	Events   []Event `json:"events"`
}

type DeviceSetEventStatusRequest struct {
	EventID []uint      `json:"event_id"`
	Status  EventStatus `json:"status"`
}
