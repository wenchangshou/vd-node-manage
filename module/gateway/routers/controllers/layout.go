package controllers

import (
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/layout"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OpenMultiScreen(c *gin.Context) {
	var service layout.OpenMultiScreenService
	if err := c.ShouldBindJSON(&service); err == nil {
		computerID, _ := strconv.Atoi(c.Param("id"))
		res := service.Open(computerID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
