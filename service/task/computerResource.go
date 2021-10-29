package task

import (
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
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

	options := make(map[string]interface{})
	options["id"] = resource.ID
	options["uri"] = "upload/" + resource.File.SourceName
	options["name"]=resource.File.Name
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


		_, err = model.AddTaskItem(task.ID, model.InstallResourceAction, options, false, depend)
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
	ID         string `json:"id" uri:"id" form:"id"`
	ResourceId string `json:"resource_id" uri:"resource_id" form:"resource_id"`
}

func (service ComputerResourceDeleteService) Delete() serializer.Response {
	// computer, err := model.GetComputerById(service.ID)
	// if err != nil {
	// 	return serializer.Err(serializer.CodeDBError, "获取计算机记录失败", err)
	// }
	// resource, err := model.GetResourceById(service.ResourceId)
	// if err != nil {
	// 	return serializer.Err(serializer.CodeDBError, "获取计算机资源失败", err)
	// }
	return serializer.Response{}

}
