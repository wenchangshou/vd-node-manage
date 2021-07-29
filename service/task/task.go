package task

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type TaskListService struct {
	Page       int               `uri:"page" json:"page" form:"page"`
	PageSize   int               `uri:"page_size" json:"page_size" form:"page_size"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}

type TaskListForm struct {
	// Items    model.Task   `json:"items"`
	model.Task
	SubItems []model.Task `json:"subItems"`
}

func (service *TaskListService) List(c *gin.Context) serializer.Response {
	conditions := make(map[string]string)
	resTaskList := make([]TaskListForm, 0)
	res, total := model.GetTasks(service.Page, service.PageSize, service.OrderBy, conditions, service.Searches)
	for _, item := range res {
		items, _ := model.GetTasksByDependID(int(item.ID))
		taskList := TaskListForm{
			Task:     item,
			SubItems: items,
		}
		resTaskList = append(resTaskList, taskList)

	}
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": resTaskList,
		},
	}
}
