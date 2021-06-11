package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
	})
}
