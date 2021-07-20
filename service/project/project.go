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
	Page       uint              `uri:"page" json:"page" form:"page"`
	PageSize   uint              `uri:"page_size" json:"page_size" form:"page_size"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}
type ProjectDetailService struct {
	ID uint `form:"path" uri:"id"`
}

func (service *ProjectDetailService) Get() serializer.Response {

	projectReleaseList, err := model.GetProjectReleaseListByProjectID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取项目版本列表失败", err)
	}
	return serializer.Response{
		Data: projectReleaseList,
	}
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

	res, total := model.GetProjects(int(service.Page), int(service.PageSize), service.OrderBy, service.Conditions, service.Searches)

	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": res,
		},
	}
}
