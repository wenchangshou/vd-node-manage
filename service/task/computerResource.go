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
	var (
		depend int
	)

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
		_resource, err := model.GetComputerResourceByComputerIdAndResourceId(int(computer), int(service.ID))
		if _resource.ID > 0 || err != nil {
			_options := make(map[string]interface{})
			_options["ID"] = _resource.ID
			_options["File"] = resource.File
			task, err := model.AddTaskItem(task.ID, model.DeleteResource, _options, false, 0)
			if err != nil {
				return serializer.Err(serializer.CodeDBError, "添加删除任务失败", err)
			}
			depend = int(task.ID)
		}
		_, err = model.AddTaskItem(task.ID, model.InstallResourceAction, options, false, uint(depend))
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加任务项失败", err)
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
