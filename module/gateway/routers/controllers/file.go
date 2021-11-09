package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/file"
)

func Upload(c *gin.Context) {
	var service file.CreateService
	if err := c.Bind(&service); err == nil {
		res := service.Create(c, CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

func DownloadFile(c *gin.Context) {
	var service file.FileDownloadService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Download(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
