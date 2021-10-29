package computer

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ListComputerResourceService struct {
	ID       string            `json:"id" uri:"id" form:"id"`
	Page     uint              `uri:"page" json:"page" form:"page"`
	PageSize uint              `uri:"page_size" json:"page_size" form:"page_size"`
	OrderBy  string            `json:"order_by" form:"order_by"`
	Category string            `form:"category" json:"category"`
	Searches map[string]string `form:"searches"`
}

type ListComputerResourceServiceResultItem struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	ID       string `json:"id"`
}

func (service *ListComputerResourceService) List() serializer.Response {
	// res := make([]ListComputerResourceServiceResultItem, 0)
	// conditions := make(map[string]string)
	// conditions["computer_id"] = service.ID
	// items, total := model.GetComputerResourceByComputerIDAndResourceCategory(int(service.Page), int(service.PageSize), service.ID, service.Category)

	// for _, item := range items {
	// 	res = append(res, ListComputerResourceServiceResultItem{
	// 		ID:       service.ID,
	// 		Name:     item.Name,
	// 		Category: item.Category,
	// 	})
	// }
	// return serializer.Response{
	// 	Data: map[string]interface{}{
	// 		"total": total,
	// 		"items": res,
	// 	},
	// }
	return serializer.Response{}

}

type ComputerResourceService struct {
	ID         string `json:"id" uri:"id" form:"id"`
	ResourceID string `json:"resource_id" form:"resource_id" uri:"resource_id"`
}

func (service ComputerResourceService) Delete() serializer.Response {
	computer, err := model.GetComputerById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机失败", err)
	}
	resource, err := model.GetResourceById(service.ResourceID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取资源失败", err)

	}
	err = computer.DeleteResource(*resource)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "删除计算机资源失败", err)
	}
	return serializer.Response{}
}
func (service ComputerResourceService) Add() serializer.Response {
	computer, err := model.GetComputerById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机失败", err)
	}
	resource, err := model.GetResourceById(service.ResourceID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取资源失败", err)
	}
	err = computer.AddResource(*resource)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "添加计算机资源失败", err)
	}
	return serializer.Response{}
}
