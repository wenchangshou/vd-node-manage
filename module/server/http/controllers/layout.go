package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/service"
	"io/ioutil"
)

func GetDeviceLayout(c *gin.Context) {
	s := service.DeviceLayoutGetService{}
	if err := c.ShouldBindUri(&s); err == nil {
		res := s.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func GetDeviceLayoutWindow(c *gin.Context) {
	s := service.DeviceLayoutGetService{}
	if err := c.ShouldBindUri(&s); err == nil {
		wid := c.Param("wid")
		s.Wid = wid
		res := s.GetWindow()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func OpenDeviceLayout(c *gin.Context) {
	s := service.DeviceLayoutOpenService{}
	if err := c.ShouldBindJSON(&s); err == nil {
		s.ID = cast.ToUint(c.Param("id"))
		s.LayoutID = cast.ToString(c.Param("layout_id"))
		res := s.Open()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func ControlLayout(c *gin.Context) {
	s := service.DeviceLayoutControlService{}
	if err := c.ShouldBindUri(&s); err == nil {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(200, serializer.ErrorResponse(err))
			return
		}
		s.Body = string(b)
		res := s.Control()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}

func CloseDeviceLayout(c *gin.Context) {
	s := service.DeviceLayoutCloseService{}
	if err := c.ShouldBindUri(&s); err == nil {
		res := s.Close()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}

}

func OpenDeviceLayoutWindow(c *gin.Context) {
	s := service.DeviceLayoutOpenWindowService{}
	if err := c.ShouldBindJSON(&s); err == nil {
		res := s.Open()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
