package resources

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/file"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/task"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/zutil"
)

type ResourceListService struct {
	Page       uint              `uri:"page" json:"page" form:"page" `
	PageSize   uint              `uri:"page_size" json:"page_size" form:"page_size"`
	OrderBy    string            `json:"order_by" form:"order_by"`
	Conditions map[string]string `form:"conditions"`
	Searches   map[string]string `form:"searches"`
}
type ResourceCreateService struct {
	Name     string `json:"name" form:"name"`
	Category string `json:"category" form:"category"`
	FileID   string `json:"file_id" form:"file_id" binding:"required" `
}

// ComputerInfo 计算机信息
type ComputerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ResourceInfo struct {
}
type ResourceListItemForm struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Category  string         `gorm:"category" json:"category"`
	Computers []ComputerInfo `json:"computers"`
}
type ResourceListForm []ResourceListItemForm

func (service ResourceListService) List() serializer.Response {
	var result ResourceListForm
	resources, total := model.ListResource(int(service.Page), int(service.PageSize), service.OrderBy, service.Conditions, service.Searches)
	for _, resource := range resources {
		computers := make([]ComputerInfo, 0)
		item := ResourceListItemForm{
			ID:       resource.ID,
			Name:     resource.Name,
			Category: resource.Category,
		}
		for _, computer := range resource.Computers {
			computers = append(computers, ComputerInfo{ID: computer.ID, Name: computer.Name})
		}
		item.Computers = computers
		result = append(result, item)
	}

	return serializer.Response{Data: map[string]interface{}{

		"total": total,
		"items": result,
	}}
}
func (service *ResourceCreateService) Create() serializer.Response {
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
	ID string `json:"id" uri:"id" form:"id"`
}

func (service ResourceDeleteService) Delete() serializer.Response {
	resource, err := model.GetResourceById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取资源失败", err)
	}
	// 如果当前的资源
	if len(resource.Computers) > 0 {
		ids := make([]string, 0)
		for _, computer := range resource.Computers {
			ids = append(ids, computer.ID)
		}
		taskItem := task.ComputerResourceDeleteService{ID: ids, Resource: resource}
		taskItem.Add()
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
	ID string `uri:"id" json:"id"`
}

// Publish 推送展项
func (service *ResourcePublishService) Publish() serializer.Response {
	clientIds := make([]string, 0)
	computers, _ := model.ListComputer()
	for _, computer := range computers {
		clientIds = append(clientIds, computer.ID)
	}
	taskItem := task.ComputerResourcePublicService{
		Computers: clientIds,
		ID:        service.ID,
	}
	return taskItem.Add()
}

type ListComputerResourceService struct {
}

type ResourceIDService struct {
	ID string `uri:"id" json:"id" form:"id"`
}

func (service ResourceIDService) Download(c *gin.Context) serializer.Response {
	resource, err := model.GetResourceById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取资源对象失败", err)
	}
	downloadService := file.FileDownloadService{
		ID: resource.FileID,
	}
	return downloadService.Download(c)
}
func (service ResourceIDService) Delete() serializer.Response {
	resource, err := model.GetResourceById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取资源对象失败", err)
	}
	fmt.Println("resource", resource)

	// tx := model.DB.Begin()
	// err = tx.Model(&model.Resource{}).Delete(resource).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return serializer.Err(serializer.CodeDBError, "删除resource记录失败", err)
	// }
	return serializer.Response{}
}
