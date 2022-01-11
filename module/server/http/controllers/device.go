package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/service"
	"io/ioutil"
	"strconv"
)

func ListDevice(c *gin.Context) {
	var (
		listService service.DeviceListService
	)
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
		if err != nil {
			c.JSON(200, serializer.ErrorResponse(errors.New("id类型错误")))
		}
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
		id := c.Param("id")
		_id, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(200, serializer.Err(serializer.CodeDBError, "id错误", err))
			return
		}
		layoutId := c.Param("layout_id")
		s.LayoutID = layoutId
		s.ID = uint(_id)
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

func GetDeviceOnline(c *gin.Context) {
	s := service.DeviceGetOnlineService{}
	if err := c.ShouldBindJSON(&s); err == nil {
		res := s.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(err))
	}

}
