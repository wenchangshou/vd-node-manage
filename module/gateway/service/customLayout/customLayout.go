package customLayout

import (
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

type Service struct {
	Name       string `json:"name"  form:"name" binding:"required"`
	Type       string `json:"type" form:"type" binding:"required"`
	Content    string `json:"content" form:"content"`
	ComputerID string
}

//Create 创建自定义布局
func (service *Service) Create() serializer.Response {

	layout := &model.CustomLayout{
		ComputerID: service.ComputerID,
		Name:       service.Name,
		Type:       service.Type,
		Content:    service.Content,
	}
	id, err := layout.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建自定义布局失败", err)
	}
	return serializer.Response{
		Data: id,
	}
}

type GetComputerCustomLayoutService struct {
	ID string `json:"id" uri:"id"`
}

func (service *GetComputerCustomLayoutService) List() serializer.Response {
	layout := model.CustomLayout{
		ComputerID: service.ID,
	}
	layouts, err := layout.GetComputerCustomLayout()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机自定义布局失败", err)
	}
	return serializer.Response{
		Data: layouts,
	}
}
