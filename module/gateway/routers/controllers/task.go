package controllers

import (
	"github.com/gin-gonic/gin"
	task2 "github.com/wenchangshou2/vd-node-manage/module/gateway/service/task"
)

// CreateProjectTask 创建项目任务
func CreateProjectTask(c *gin.Context) {
	var service task2.ComputerProject
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CreateResourceTask 创建资源任务
func CreateResourceTask(c *gin.Context) {
	var service task2.ComputerResourcePublicService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func DeleteProjectTask(c *gin.Context) {
	var service task2.ComputerProject
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListTask 当前任务列表
func ListTask(c *gin.Context) {
	var service task2.ListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UpdateTask(c *gin.Context) {
	var service task2.UpdateTaskService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func UpdateTaskItem(c *gin.Context) {
	var service task2.UpdateTaskItemService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
