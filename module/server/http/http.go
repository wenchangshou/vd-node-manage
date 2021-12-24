package http

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/wenchangshou2/vd-node-manage/module/server/g"
	"github.com/wenchangshou2/vd-node-manage/module/server/http/controllers"
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
		device.POST("/list", controllers.ListDevice)
		device.POST("", controllers.AddDevice)
		device.DELETE("/:id", controllers.DeleteDevice)
		device.GET("/:id", controllers.GetDevice)
		device.POST("/register", controllers.RegisterDevice)
		device.POST("/:id/resource", controllers.AddDeviceResource)
		device.GET("/:id/resource", controllers.ListDeviceResource)
		device.POST("/:id/layout", controllers.SetDeviceLayout)
		device.DELETE("/:id/layout", controllers.CloseDeviceLayout)
		device.POST("/:id/control", controllers.ControlLayout)
		resource := v1.Group("/resource")
		resource.POST("/upload", controllers.UploadFile)
		resource.POST("", controllers.AddResource)
	}
	return r
}
func Start() {
	fmt.Println(g.Config().Http)
	if !g.Config().Http.Enabled {
		return
	}
	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}
	r := InitRouter()
	if err := r.Run(addr); err != nil {
		log.Fatalln("start http", "err", err.Error())
	}
}
