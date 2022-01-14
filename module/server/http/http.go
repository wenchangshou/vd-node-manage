package http

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/wenchangshou/vd-node-manage/module/server/g"
	"github.com/wenchangshou/vd-node-manage/module/server/http/controllers"
)

var router *httprouter.Router

type Dto struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", controllers.Health)
		device := v1.Group("/device")
		{
			device.POST("/list", controllers.ListDevice)
			device.POST("/online", controllers.GetDeviceOnline)
			device.POST("", controllers.AddDevice)
			device.POST("/lease", controllers.SetDeviceExpired)
			device.DELETE("/:id", controllers.DeleteDevice)
			device.GET("/:id", controllers.GetDevice)
			device.POST("/register", controllers.RegisterDevice)
			device.POST("/:id/resource", controllers.AddDeviceResource)
			device.DELETE("/:id/resource/:resource_id", controllers.DeleteDeviceResource)
		}
		layout := v1.Group("/layout")
		{
			layout.POST("/:id/:layout_id", controllers.OpenDeviceLayout)
			layout.DELETE("/:id/:layout_id", controllers.CloseDeviceLayout)
			layout.GET("/:id/:layout_id", controllers.GetDeviceLayout)
			layout.GET("/:id/:layout_id/:wid", controllers.GetDeviceLayoutWindow)
			layout.POST("/:id/:layout_id/:window_id/control", controllers.ControlLayout)
		}
		window := v1.Group("/window")
		{
			window.POST("", controllers.OpenDeviceLayoutWindow)
		}
		resource := v1.Group("/resource")
		{
			resource.POST("/upload", controllers.UploadFile)
			resource.POST("", controllers.AddResource)
			resource.POST("/list", controllers.ListDeviceResource)
		}
		event := v1.Group("/event")
		{
			event.POST("/list", controllers.ListEvent)
		}
	}
	return r
}
func Start() {
	fmt.Println(g.Config().Http)
	if !g.Config().Http.Enabled {
		return
	}

	addr := g.Config().Http.Listen
	if g.Config().Mode == "docker" {
		addr = os.Getenv("LISTEN_ADDR") + ":6031"
	}
	if addr == "" {
		return
	}

	r := InitRouter()
	if err := r.Run(addr); err != nil {
		log.Fatalln("start http", "err", err.Error())
	}
}
