package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ResourceListService struct {
	Page       uint              `uri:"page" json:"page" form:"page"`
	PageSize   uint              `uri:"page_size" json:"page_size" form:"page_size"`
	OrderBy    string            `json:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}
type ResourceCreateService struct {
	Name     string `json:"name" form:"name"`
	Category string `json:"category" form:"category"`
	FileID   uint   `json:"file_id" form:"file_id"`
}

func (service *ResourceListService) List(c *gin.Context) serializer.Response {
	res, total := model.GetResources(int(service.Page), int(service.PageSize), service.OrderBy, service.Conditions, service.Searches)
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": res,
		},
	}
}
func (service *ResourceCreateService) Create(c *gin.Context) serializer.Response {
	resource := model.Resource{
		Name:     service.Name,
		Category: service.Category,
		FileID:   service.FileID,
	}
	id, err := resource.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建资源记录失败", err)
	}
	return serializer.Response{
		Data: id,
	}
}
