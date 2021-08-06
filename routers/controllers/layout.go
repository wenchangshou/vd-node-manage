package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/layout"
)

func OpenMultiScreen(c *gin.Context) {
	var service layout.LayoutOpenMultiScreenService
	if err := c.ShouldBindJSON(&service); err == nil {
		computerID, _ := strconv.Atoi(c.Param("id"))
		res := service.Open(computerID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
