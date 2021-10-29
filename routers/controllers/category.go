package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/exhibition"
)

func CreateExhibitionCategory(c *gin.Context) {
	var service exhibition.CategoryService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
