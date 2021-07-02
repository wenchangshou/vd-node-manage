package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/project"
)

// CreateProject 创建新的项目
func CreateProject(c *gin.Context) {
	var service project.ProjectCreateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create(c, CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ListProjest(c *gin.Context) {
	var service project.ProjectListService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.List(c, CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
