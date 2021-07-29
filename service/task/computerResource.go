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

func (service *ComputerResource) Add() serializer.Response {

	resource, err := model.GetResourceById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFindComputerResource, "没有找到资源", err)
	}

	options := make(map[string]interface{})
	options["ID"] = resource.ID
	options["File"] = resource.File
	options["Uri"] = "upload/" + resource.File.SourceName
	for _, computer := range service.Computers {
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
