package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/server/service"
)

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
