package exhibition

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type CategoryService struct {
	Name       string `form:"name" json:"name" binding:"required"`
	ComputerID string `form:"computer_id" json:"computer_id" binding:"required"`
}

func (service *CategoryService) Create() serializer.Response {
	category := model.ExhibitionCategory{
		Name:       service.Name,
		ComputerID: service.ComputerID,
	}
	id, err := category.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建类别失败", err)
	}
	return serializer.Response{
		Data: id,
	}
}
