package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/service"
)

func ListEvent(c *gin.Context) {
	s := service.DeviceEventGetService{}
	if err := c.BindJSON(&s); err == nil {
		res := s.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
