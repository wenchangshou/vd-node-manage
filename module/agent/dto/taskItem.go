package dto

type TaskItem struct {
	ID       string  `json:"id"`
	Options  string  `json:"options"`
	Status   int     `json:"status"`
	Depend   string  `json:"depend"`
	Schedule float32 `json:"schedule"`
	Active   bool    `json:"active"`
	Action int `json:"action"`
}
