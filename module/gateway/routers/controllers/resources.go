package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/resources"
)

func ListResource(c *gin.Context) {
	var service resources.ResourceListService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func DeleteResource(c *gin.Context) {
	var service resources.ResourceDeleteService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func CreateResource(c *gin.Context) {
	var service resources.ResourceCreateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func PublishResource(c *gin.Context) {
	var service resources.ResourcePublishService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Publish()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func DownloadResourceFile(c *gin.Context) {
	var service resources.ResourceIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Download(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func DeleteResourceFile(c *gin.Context) {
	var service resources.ResourceIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}