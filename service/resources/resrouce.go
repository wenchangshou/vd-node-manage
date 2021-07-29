package resources

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
	"github.com/wenchangshou2/vd-node-manage/service/task"
	"github.com/wenchangshou2/zutil"
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
	FileID   uint   `json:"file_id" form:"file_id" binding:"required" `
}
type ResourceListItemForm struct {
	model.Resource
	Computers []string
}
type ResourceListForm []ResourceListItemForm

func (service *ResourceListForm) AppendComputet(resourceID int, hostName string) {
	for k, v := range *service {
		if v.Resource.ID == uint(resourceID) {
			(*service)[k].Computers = append((*service)[k].Computers, hostName)
		}
	}
}
func (service *ResourceListService) List(c *gin.Context) serializer.Response {
	var result ResourceListForm
	ids := make([]int, 0)
	res, total := model.GetResources(int(service.Page), int(service.PageSize), service.OrderBy, service.Conditions, service.Searches)
	for _, item := range res {
		ids = append(ids, int(item.ID))
	}
	computerResourceList, err := model.GetComputerResourcesByIds(ids)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机资源失败", err)
	}
	for _, resource := range res {
		item := ResourceListItemForm{
			Computers: make([]string, 0),
		}
		item.Resource = resource
		result = append(result, item)
	}
	computers, _ := model.ListComputer()
	for _, computerResource := range computerResourceList {
		for _, computer := range computers {
			if computer.ID == computerResource.ComputerID {
				result.AppendComputet(int(computerResource.ResourceID), computer.Name)
			}
		}
	}

	fmt.Println("ids", ids)
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": result,
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

type ResourceDeleteService struct {
	ID uint `json:"id" uri:"id" form:"id"`
}

func (service ResourceDeleteService) Delete(c *gin.Context) serializer.Response {
	resource, err := model.GetResourceById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取资源失败", err)
	}
	resourcePath := path.Join("upload", resource.File.SourceName)
	err = zutil.IsExistDelete(resourcePath)
	if err != nil {
		return serializer.Err(serializer.CodeFileDeleteErr, "删除文件失败", err)
	}
	err = resource.File.Delete()
	if err != nil {
		return serializer.Err(serializer.CodeDeleteFileRecordErr, "删除文件记录失败", err)
	}
	err = resource.Delete()
	if err != nil {
		return serializer.Err(serializer.CodeDeleteResourceRecordErr, "删除资源记录失败", err)
	}
	return serializer.Response{}

}

type ResourcePublishService struct {
	ID uint `uri:"id" json:"id"`
}

func (service *ResourcePublishService) Publish() serializer.Response {
	clientsids := make([]uint, 0)
	computers, _ := model.ListComputer()
	for _, computer := range computers {
		clientsids = append(clientsids, computer.ID)
	}
	taskItem := task.ComputerResource{
		Computers: clientsids,
		ID:        service.ID,
	}
	return taskItem.Add()
}
