package task

import (
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ComputerResource struct {
	Computers []uint `json:"computers" binding:"required"`
	ID        uint   `json:"id"`
}
type installResourceOption struct {
	ID   uint `json:"id"`
	File model.File
}

func (server *ComputerResource) Add() serializer.Response {
	computerResource, err := model.GetComputerResourceById(int(server.ID))
	if computerResource == nil || err != nil {
		return serializer.Err(serializer.CodeNotFindComputerResource, "没有找到计算机资源", err)
	}
	resource, err := model.GetResourceById(computerResource.ResourceID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFindComputerResource, "没有找到资源", err)
	}

	options := make(map[string]interface{})
	options["ID"] = computerResource.ID
	options["File"] = resource.File
	for _, computer := range server.Computers {
		task, err := model.AddTask(fmt.Sprintf("添加%s资源", resource.Name), computer)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
		}
		_, err = model.AddTaskItem(task.ID, model.InstallResourceAction, options, false, 0)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加任务项失败", err)
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
