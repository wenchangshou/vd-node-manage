package module

import (
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

type CreateModuleService struct {
	Name     string `json:"name" form:"name"`
	Category string `json:"category" form:"category"`
	Value    string `json:"value" form:"value"`
}

func (service CreateModuleService) Create() serializer.Response {
	module := model.Module{
		Name:     service.Name,
		Value:    service.Value,
		Category: service.Category,
	}
	err := module.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建模块失败", err)
	}
	return serializer.Response{
		Data: module.ID,
	}
}

type DeleteModuleService struct {
	ID string `json:"id" uri:"id"`
}

func (service *DeleteModuleService) Delete() serializer.Response {
	ids, err := model.GetSpecifiedModuleExhibition(service.ID)
	if len(ids) > 0 {
		return serializer.Err(serializer.CodeDBError, "当前的模块正在被展项所使用", err)
	}
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "删除模块失败", err)
	}
	err = model.DeleteModuleByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "删除模块失败", err)
	}
	return serializer.Response{}
}

type ListModuleService struct {
	Page       int               `json:"page" binding:"min=1,required"`
	PageSize   int               `json:"page_size" binding:"min=1,required"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}
type ListModuleItemForm struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
type ListModuleForm []ListModuleItemForm

func (service *ListModuleService) List() serializer.Response {
	var result ListModuleForm
	res, total := model.ListModules(int(service.Page), int(service.PageSize), service.OrderBy, service.Conditions, service.Searches)
	for _, module := range res {
		item := ListModuleItemForm{
			ID:    module.ID,
			Name:  module.Name,
			Value: module.Value,
		}
		result = append(result, item)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": result,
		},
	}
}
