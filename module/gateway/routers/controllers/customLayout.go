package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/customLayout"
)

func CreateCustomLayout(c *gin.Context) {
	var service customLayout.Service
	if err := c.ShouldBindJSON(&service); err == nil {
		computerID := c.Param("id")
		service.ComputerID = computerID
		result := service.Create()
		c.JSON(200, result)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetComputerCustomLayout(c *gin.Context) {
	var service customLayout.GetComputerCustomLayoutService
	if err := c.ShouldBindUri(&service); err == nil {
		result := service.List()
		c.JSON(200, result)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
