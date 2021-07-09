package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/computer"
)

func UpdateComputer(c *gin.Context) {
	var service computer.ComputerUpdateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func ListComputer(c *gin.Context) {
	var service computer.ComputerListService
	res := service.List(c)
	c.JSON(200, res)
}
