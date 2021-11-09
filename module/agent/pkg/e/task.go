package e

// TaskItem 任务项
type TaskItem struct {
	ID      string `json:"id"`
	Action  int    `json:"action"`
	Status  int    `json:"status"`
	Options string `json:"options"`
	Depend  string `json:"depend"`
}

// TaskGroup 任务组
type TaskGroup struct {
	Name     string     `json:"name"`
	ID       string     `json:"id"`
	TaskList []TaskItem `json:"task_list"`
}
