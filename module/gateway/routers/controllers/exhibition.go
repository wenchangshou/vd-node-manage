package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/exhibition"
)

func GetComputerExhibition(c *gin.Context) {
	var service exhibition.GetComputerExhibitionService
	if err := c.BindUri(&service); err == nil {
		res := service.Get(c)
		c.JSON(200, res)
	}
}
func CreateComputerExhibition(c *gin.Context) {
	var service exhibition.ExhibitionService
	if err := c.BindJSON(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetExhibition(c *gin.Context) {
	var service exhibition.GetExhibitionDetailsService
	if err := c.BindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

func UpdateExhbition(c *gin.Context) {
	var service exhibition.ExhibitionService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

func DeleteExhibition(c *gin.Context) {
	var service exhibition.ExhibitionService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
