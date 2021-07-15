package project

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/hashid"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ProjectCreateService struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Category    string `form:"category" json:"category" binding:"required"`
	Description string `json:"description"`
	Arguments   string `json:"arguments"`
}

type ProjectListService struct {
	Page       int               `json:"page" binding:"min=1,required"`
	PageSize   int               `json:"page_size" binding:"min=1,required"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}

// Create 创建一个新的项目
func (service *ProjectCreateService) Create(c *gin.Context, user *model.User) serializer.Response {
	project := model.Project{
		Name:        service.Name,
		Category:    service.Category,
		Description: service.Description,
		Arguments:   service.Arguments,
	}
	id, err := project.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建项目失败", err)
	}
	return serializer.Response{
		Data: hashid.HashID(id, hashid.ProjectID),
	}
}
func (service *ProjectListService) List(c *gin.Context, user *model.User) serializer.Response {
	res, total := model.GetProjects(service.Page, service.PageSize, service.OrderBy, service.Conditions, service.Searches)

	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": res,
		},
	}
}
