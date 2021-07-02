package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/project"
)

// CreateProjectRelease 创建项目发行版本
func CreateProjectRelease(c *gin.Context) {
	var service project.ProjectReleaseCreateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create(c, CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetProjectRelease 获取单个发行版本
func GetProjectRelease(c *gin.Context) {
	var service project.GetProjectReleaseService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}