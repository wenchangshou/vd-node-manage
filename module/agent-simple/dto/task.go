package dto

type Task struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Items []TaskItem `json:"items"`
}
