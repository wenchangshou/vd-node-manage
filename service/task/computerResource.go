package task

import (
	"encoding/json"

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
	computerResource, err := model.GetComputerResource(server.ID)
	if computerResource == nil || err != nil {
		return serializer.Err(serializer.CodeNotFindComputerResource, "没有找到计算机资源", err)
	}
	resource, err := model.GetResourceById(computerResource.ResourceId)
	if err != nil {
		return serializer.Err(serializer.CodeNotFindComputerResource, "没有找到资源", err)
	}

	options := installResourceOption{
		ID:   computerResource.ID,
		File: resource.File,
	}
	optionByte, _ := json.Marshal(options)
	for _, computer := range server.Computers {
		task := model.NewTask(int(computer), model.InstallResourceAction, string(optionByte), 0)
		task.Add()
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
