package project

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

type CreateService struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Category    string `form:"category" json:"category" binding:"required"`
	Description string `json:"description"`
	Arguments   string `json:"arguments"`
	Start       string `json:"start"`
	File        string `json:"file"`
}

type ListService struct {
	Page       uint              `uri:"page" json:"page" form:"page"`
	PageSize   uint              `uri:"page_size" json:"page_size" form:"page_size"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}
type DetailService struct {
	ID string `form:"path" uri:"id"`
}
type ListItemForm struct {
	model.Project
	Computers []string
}
type ListForm []ListItemForm

func (service *ListForm) AppendComputet(projectId string, hostName string) {
	for k, v := range *service {
		if v.Project.ID == projectId {
			(*service)[k].Computers = append((*service)[k].Computers, hostName)
		}
	}
}

func (service *DetailService) Get() serializer.Response {

	projectReleaseList, err := model.GetProjectReleaseListByProjectID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取项目版本列表失败", err)
	}
	return serializer.Response{
		Data: projectReleaseList,
	}
}

// Create 创建一个新的项目
func (service *CreateService) Create(c *gin.Context, user *model.User) serializer.Response {
	project := model.Project{
		Name:        service.Name,
		Category:    service.Category,
		Description: service.Description,
		Arguments:   service.Arguments,
		Start:       service.Start,
		Cover:       service.File,
	}
	id, err := project.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建项目失败", err)
	}
	return serializer.Response{
		Data: id,
	}
}
func (service *ListService) List(c *gin.Context, user *model.User) serializer.Response {
	var result ListForm
	queryCondition := make([]string, 0)
	res, total := model.GetProjects(int(service.Page), int(service.PageSize), service.OrderBy, service.Conditions, service.Searches)
	for _, project := range res {
		item := ListItemForm{
			Computers: make([]string, 0),
		}
		item.Project = project
		result = append(result, item)
		queryCondition = append(queryCondition, project.ID)
	}
	// computerProjectList, err := model.GetComputerProjectByProjectIds(queryCondition)
	// if err != nil {
	// 	return serializer.Err(serializer.CodeDBError, "获取计算机项目资源列表失败", err)
	// }
	// computers, _ := model.ListComputer()
	// for _, computerProject := range computerProjectList {
	// 	for _, computer := range computers {
	// 		if computer.ID == computerProject.ComputerId {
	// 			result.AppendComputet(computerProject.ProjectID, computer.Name)
	// 		}
	// 	}
	// }

	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": result,
		},
	}
}
