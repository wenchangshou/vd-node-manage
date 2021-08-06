package controllers

import (
	"fmt"
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
	var service computer.ComputerProjectListService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

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
func OpenComputerExhibition(c *gin.Context) {
	var service computer.ComputerExhibitionOpenService
	if err := c.ShouldBindUri(service); err == nil {
	}
}
func GetComputerProjectDir(c *gin.Context) {
	computerID := c.Param("id")
	projectID := c.Param("projectID")
	fmt.Printf("computerID:%v,projectID:%v", computerID, projectID)
	var service computer.ComputerProjectDirectoryService
	if err := c.ShouldBindJSON(&service); err == nil {
		service.ComputerID, _ = strconv.Atoi(computerID)
		service.ProjectID, _ = strconv.Atoi(projectID)
		service.Get()
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
