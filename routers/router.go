package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/routers/controllers"
)

// InitCORS 初始化跨域配置
func InitCORS(router *gin.Engine) {
	if conf.CORSConfig.AllowOrigins[0] != "UNSET" {
		if conf.CORSConfig.AllowOrigins[0] != "UNSET" {
			router.Use(cors.New(cors.Config{
				AllowOrigins:     conf.CORSConfig.AllowOrigins,
				AllowMethods:     conf.CORSConfig.AllowMethods,
				AllowHeaders:     conf.CORSConfig.AllowHeaders,
				AllowCredentials: conf.CORSConfig.AllowCredentials,
				ExposeHeaders:    conf.CORSConfig.ExposeHeaders,
			}))
			return
		}
		return
	}

}
func InitMasterRouter() *gin.Engine {
	r := gin.Default()
	InitCORS(r)
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
	v1 := r.Group("/api/v1")
	{
		site := v1.Group("site")
		{
			site.GET("ping", controllers.Ping)
		}
	}
	return r
}
func InitRouter() *gin.Engine {
	return InitMasterRouter()
}
