package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/task"
)

func CreateTask(c *gin.Context) {
	var service task.AddTaskItemService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// QueryUserTask 查询用户任务
func QueryUserPendingTask(c *gin.Context) {
	var service task.QueryComputerTaskService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Query()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
