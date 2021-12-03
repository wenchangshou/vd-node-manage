package dto

type Task struct {
	Action string `json:"action"`
	Params map[string]interface{}
}
