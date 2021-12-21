package model

type EventRequest struct {
	EventID   string `json:"event_id"`
	Action    string `json:"action"`
	DeviceID  uint   `json:"device_id"`
	Arguments []byte `json:"arguments"`
	Reply     bool   `json:"reply"`
}
type EventReply struct {
	EventID string `json:"event_id"`
	Err     error  `json:"err"`
	Msg     string `json:"msg"`
	Body    string `json:"body"`
}

func GenerateSimpleSuccessEventReply(id string) EventReply {
	reply := EventReply{
		EventID: id,
		Err:     nil,
		Body:    "",
	}
	return reply
}

// OpenLayoutCmdParams 打开布局命令参数
type OpenLayoutCmdParams struct {
	ID      string                 `json:"id"`
	Kill    bool                   `json:"kill"`
	Style   map[string]interface{} `json:"style"`
	Windows []OpenWindowInfo       `json:"windows"`
}
