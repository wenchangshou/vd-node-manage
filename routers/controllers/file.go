package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/file"
)

func Upload(c *gin.Context) {
	var service file.FileCreateService
	if err := c.Bind(&service); err == nil {
		res := service.Create(c, CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

	// _file, err := c.FormFile("file")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, ErrorResponse(err))
	// 	return
	// }
	// ext := filepath.Ext(_file.Filename)
	// newFileName := uuid.New().String() + ext
	// sourceName := "upload/" + newFileName
	// if err := c.SaveUploadedFile(_file, sourceName); err != nil {
	// 	c.JSON(http.StatusBadRequest, ErrorResponse(err))
	// 	return
	// }
	// service := &file.FileCreateService{
	// 	Name:       _file.Filename,
	// 	SourceName: sourceName,
	// 	UserId:     CurrentUser(c).ID,
	// 	Size:       uint(_file.Size),
	// }
	// res := service.Create(c, CurrentUser(c))
	// c.JSON(200, res)

}
