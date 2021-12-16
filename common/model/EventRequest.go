package model

type EventRequest struct {
	Action    string `json:"action"`
	DeviceID  uint   `json:"device_id"`
	Arguments []byte `json:"arguments"`
}

// OpenLayoutCmdParams 打开布局命令参数
type OpenLayoutCmdParams struct {
	ID      string                 `json:"id"`
	Kill    bool                   `json:"kill"`
	Style   map[string]interface{} `json:"style"`
	Windows []OpenWindowInfo       `json:"windows"`
}
