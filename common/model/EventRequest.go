package model

type EventRequest struct {
	Action    string      `json:"action"`
	Arguments interface{} `json:"arguments"`
}

type OpenLayoutCmdParams struct {
	ID      string                 `json:"id"`
	Kill    bool                   `json:"kill"`
	Style   map[string]interface{} `json:"style"`
	Windows []OpenWindowInfo       `json:"windows"`
}
