package controllers

import (
	"strconv"

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
func UpdateComputerName(c *gin.Context) {
	var service computer.ComputerUpdateNameService
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	id := c.Param("id")
	service.ID, _ = strconv.Atoi(id)
	res := service.Update(c)
	c.JSON(200, res)
}
func ListComputer(c *gin.Context) {
	var service computer.ComputerListService
	res := service.List(c)
	c.JSON(200, res)
}
func ListComputerProject(c *gin.Context) {

}

func GetComputerDetails(c *gin.Context) {
	var service computer.ComputerGetDetailsService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func GetCrossResources(c *gin.Context) {
	var service computer.ComputerProjectGetCrossResource
	res := service.Get()
	c.JSON(200, res)

}
