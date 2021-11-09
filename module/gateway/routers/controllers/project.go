package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/project"
)

// CreateProject 创建新的项目
func CreateProject(c *gin.Context) {
	var service project.CreateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create(c, CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ListProject(c *gin.Context) {
	var service project.ListService
	if err := c.Bind(&service); err == nil {
		res := service.List(c, CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func GetProjectReleaseList(c *gin.Context) {
	var service project.DetailService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
