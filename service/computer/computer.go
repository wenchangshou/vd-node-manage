package computer

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ComputerListService struct {
}

func (service *ComputerListService) List(c *gin.Context) serializer.Response {
	computer, total := model.ListComputer()
	return serializer.Response{
		Data: map[string]interface{}{
			"total": total,
			"items": computer,
		},
	}
}

type ComputerUpdateService struct {
	Ip       string `json:"ip" form:"ip" binding:"required"`
	Mac      string `json:"mac" form:"mac" binding:"required"`
	HostName string `json:"host_name"`
}

func (service *ComputerUpdateService) Update(c *gin.Context) serializer.Response {
	client := model.Computer{
		Ip:       service.Ip,
		Mac:      service.Mac,
		HostName: service.HostName,
	}
	if client.IsExistByMac() {
		if err := client.UpdateByMac(); err != nil {
			return serializer.Err(403, "更新客户端信息失败", err)
		}
		return serializer.Response{}
	}
	if err := client.Create(); err != nil {
		return serializer.Err(403, "创建客户端信息失败", err)
	}
	return serializer.Response{}
}

type ComputerGetDetailsService struct {
	ID int `json:"id" form:"id" uri:"id"`
}
type ComputerGetDetailsForm struct {
	model.Computer
	Projects  []model.ProjectRelease
	Resources []model.ComputerResource
}

func (service *ComputerGetDetailsService) Get(c *gin.Context) serializer.Response {
	computerForm := ComputerGetDetailsForm{}
	computer, err := model.GetComputerById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机信息失败", err)
	}
	computerForm.Computer = computer
	projects, err := model.GetComputerProjectByComputerID(int(computer.ID))
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机项目列表失败", err)
	}
	computerForm.Projects = projects
	resources, err := model.GetComputerResourceByComputerId(int(computer.ID))
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机资源失败", err)
	}
	computerForm.Resources = resources
	return serializer.Response{Data: map[string]interface{}{
		"items": computerForm,
	}}
}

type ComputerUpdateNameService struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}

func (service *ComputerUpdateNameService) Update(c *gin.Context) serializer.Response {
	data := make(map[string]interface{})
	data["name"] = service.Name
	err := model.UpdateComputerById(service.ID, data)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "设置计算机信息失败", err)
	}
	return serializer.Response{}
}
