package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
	"github.com/wenchangshou2/vd-node-manage/service/system"
)

func GetExtranet(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Data: c.ClientIP(),
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
