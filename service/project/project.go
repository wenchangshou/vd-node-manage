package project

import (
	"strings"

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
	var res []model.Project
	var total int64 = 0
	tx := model.DB.Model(&model.Project{})
	if service.OrderBy != "" {
		tx = tx.Order(service.OrderBy)
	}
	for k, v := range service.Conditions {
		tx = tx.Where(k+" = ?", v)
	}
	if len(service.Searches) > 0 {
		search := ""
		for k, v := range service.Searches {
			search += (k + " like '%" + v + "%' OR ")
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize).Find((&res))
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": res,
		},
	}
}
