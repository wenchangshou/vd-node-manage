package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/service"
)

func AddResource(c *gin.Context) {
	service := &service.ResourceAddService{}
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func UploadFile(_ *gin.Context) {

}
func DeleteDeviceResource(c *gin.Context) {
	service := &service.DeviceResourceDeleteService{}
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
