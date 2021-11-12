package task

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

type ComputerResourcePublicService struct {
	Computers []string `json:"computers" binding:"required"`
	ID        string   `json:"id"`
}

func (service *ComputerResourcePublicService) EqualsResource(computers []*model.Computer, id string) bool {
	for _, computer := range computers {
		if computer.ID == id {
			return true
		}
	}
	return false
}

// Add 添加新的展项
func (service *ComputerResourcePublicService) Add() serializer.Response {
	var (
		depend string
	)

	resource, err := model.GetResourceById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFindComputerResource, "没有找到资源", err)
	}

	option := make(map[string]interface{})
	option["id"] = resource.ID
	option["uri"] = "upload/" + resource.File.SourceName
	option["name"] = resource.File.Name
	option["source"] = resource.File.SourceName
	for _, computer := range service.Computers {
		task, err := model.AddTask(fmt.Sprintf("添加%s资源", resource.Name), computer)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
		}
		if service.EqualsResource(resource.Computers, computer) {
			_options := make(map[string]interface{})
			_options["file"] = resource.File
			task, err := model.AddTaskItem(task.ID, model.DeleteResource, _options, false, "")
			if err != nil {
				return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
			}
			depend = task.ID
		}

		_, err = model.AddTaskItem(task.ID, model.InstallResourceAction, option, false, depend)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加任务项失败", err)
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

type ComputerResourceDeleteService struct {
	ID       []string        `json:"id" uri:"id" form:"id"`
	Resource *model.Resource `json:"resource"`
}

func (service ComputerResourceDeleteService) Add() serializer.Response {
	options := make(map[string]interface{})
	options["id"] = service.Resource.ID
	options["name"] = service.Resource.ID + "-" + service.Resource.Name
	for _, computer := range service.ID {
		task, err := model.AddTask(fmt.Sprintf("删除%s资源", service.Resource.Name), computer)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "删除任务失败", err)
		}
		_, err = model.AddTaskItem(task.ID, model.DeleteResource, options, false, "")
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加任务项失败", err)
		}
	}

	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}

}
