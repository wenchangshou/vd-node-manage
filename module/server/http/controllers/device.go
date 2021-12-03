package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/service"
)

func ListDevice(c *gin.Context) {
	var (
		listService service.DeviceListService
	)
	if err := c.ShouldBindJSON(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func AddDevice(c *gin.Context) {
	ser := &service.DeviceCreateService{}
	if err := c.BindJSON(&ser); err == nil {
		res := ser.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func RegisterDevice(c *gin.Context) {
	ser := &service.DeviceRegisterService{}
	if err := c.BindJSON(&ser); err == nil {
		res, err := ser.Register()
		if err != nil {
			c.JSON(200, serializer.ErrorResponse(err))
			return
		}
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}

}

// AddDeviceResource 添加设备资源
func AddDeviceResource(c *gin.Context) {
	s := &service.DeviceResourceAddService{}
	if err := c.BindJSON(&s); err == nil {
		res := s.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
