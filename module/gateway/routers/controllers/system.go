package controllers

import (
	"bytes"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/service/system"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExtranet(c *gin.Context) {
	ip := c.ClientIP()
	c.JSON(200, serializer.Response{
		Code: 0,
		Data: ip,
	})
}
func ExportProjectRecord(c *gin.Context) {
	var service system.ProjectExportService
	f, err := service.Export()
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	buffer := new(bytes.Buffer)
	f.Write(buffer)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "project.xlsx"))
	c.Writer.Header().Set("Content-Type", "application/*")
	c.Writer.Write(buffer.Bytes())
	fmt.Println(f, err)
}
