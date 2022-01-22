package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cast"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/service"
)

func ListDevice(c *gin.Context) {
	var (
		listService service.DeviceListService
	)
	span, _ := opentracing.StartSpanFromContext(c, "span_foo3")
	defer func() {
		//4.接口调用完，在tag中设置request和reply
		span.SetTag("request", c.Request)
		span.Finish()
	}()
	if err := c.ShouldBindJSON(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func GetDevice(c *gin.Context) {
	var (
		service service.DeviceGetService
	)
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))

	}
}

func SetDeviceExpired(c *gin.Context) {
	s := service.UpdateDeviceStruct{}
	if err := c.BindJSON(&s); err == nil {
		res := s.SetLease()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func AddDevice(c *gin.Context) {
	s := &service.DeviceCreateService{}
	if err := c.BindJSON(&s); err == nil {
		res := s.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func DeleteDevice(c *gin.Context) {
	s := &service.DeviceDeleteService{}
	if err := c.ShouldBindUri(&s); err == nil {
		res := s.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func RegisterDevice(c *gin.Context) {
	ser := &service.DeviceRegisterService{}
	if err := c.BindJSON(&ser); err == nil {
		res, err := ser.Register()
		if err != nil {
			c.JSON(200, serializer.ErrorResponse(err))
			return
		}
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}

}

// AddDeviceResource 添加设备资源
func AddDeviceResource(c *gin.Context) {
	s := &service.DeviceResourceAddService{}
	if err := c.BindJSON(&s); err == nil {
		s.ID = cast.ToUint(c.Param("id"))
		res := s.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}
func AddDeviceProject(c *gin.Context) {
	s := &service.DeviceProjectAddService{}
	if err := c.BindJSON(&s); err == nil {
		s.ID = cast.ToUint(c.Param("id"))
		res := s.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}
}

func ListDeviceResource(c *gin.Context) {
	s := service.DeviceResourceListService{}
	if err := c.ShouldBindJSON(&s); err == nil {
		res := s.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}

}
func GetDeviceOnline(c *gin.Context) {
	s := service.DeviceGetOnlineService{}
	if err := c.ShouldBindJSON(&s); err == nil {
		res := s.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}

}
