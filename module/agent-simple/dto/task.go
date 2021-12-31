package dto

import "github.com/wenchangshou/vd-node-manage/common/model"

//Task 任务栏
type Task struct {
	Action model.EventStatus `json:"action"`
	ID     uint              `json:"id"`
	Params map[string]interface{}
	Status model.EventStatus `json:"status"`
}
