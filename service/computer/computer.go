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
