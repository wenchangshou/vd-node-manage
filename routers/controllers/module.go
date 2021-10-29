package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/module"
)

func ListModule(c *gin.Context) {
	var service module.ListModuleService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
func CreateModule(c *gin.Context) {
	var service module.CreateModuleService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

func DeleteModule(c *gin.Context) {

}
func UpdateModule(c *gin.Context) {

}
