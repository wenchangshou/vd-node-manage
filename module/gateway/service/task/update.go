package task

import (
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/serializer"
)

type UpdateTaskService struct{
	ID []string `json:"id"`
	Status int `json:"status"`
}
func (s UpdateTaskService)Update() serializer.Response {
	err:=model.UpdateTaskStatusByIds(s.ID,s.Status)
	if err!=nil{
		return serializer.Err(serializer.CodeDBError,"更新任务状态失败",err)
	}
	return serializer.Response{}
}
type UpdateTaskItemService struct{
	ID []string `json:"id"`
	Status int `json:"status"`
}
func (s *UpdateTaskItemService)Update() serializer.Response {
	err:=model.UpdateTaskItemStatusByIds(s.ID,s.Status)
	if err!=nil{
		return serializer.Err(serializer.CodeDBError,"更新任务状态失败",err)
	}
	return serializer.Response{}
}