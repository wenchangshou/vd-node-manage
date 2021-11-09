package task

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/serializer"
)

type ListService struct {
	Page       int               `uri:"page" json:"page" form:"page"`
	PageSize   int               `uri:"page_size" json:"page_size" form:"page_size"`
	OrderBy    string            `json:"order_by" form:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}

type ListForm struct {
	// Items    model.Task   `json:"items"`
	model.Task
	SubItems []model.Task `json:"subItems"`
}

func (service *ListService) List(c *gin.Context) serializer.Response {
	res, total := model.ListTask(service.Page, service.PageSize, service.OrderBy, service.Conditions, service.Searches)
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": res,
		},
	}
}
