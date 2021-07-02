package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/middleware"
	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/routers/controllers"
)

// InitCORS 初始化跨域配置
func InitCORS(router *gin.Engine) {
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

}
func InitMasterRouter() *gin.Engine {
	r := gin.Default()
	InitCORS(r)
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Session(conf.SystemConfig.SessionSecret))
	v1.Use(middleware.CurrentUser())
	{
		site := v1.Group("site")
		{
			site.GET("ping", controllers.Ping)
		}
		user := v1.Group("user")
		{
			user.POST("session", controllers.UserLogin)
			user.POST("",
				controllers.UserRegister)
		}
		system := v1.Group("system")
		{
			system.GET("extranet", controllers.GetExtranet)
		}
		client := v1.Group("computer")
		{
			client.PUT("", controllers.UpdateComputer)
			client.GET("/:mac/task/pending", controllers.QueryUserPendingTask)
		}

		projectRelease := v1.Group("projectRelease")
		projectRelease.GET(":id", controllers.GetProjectRelease)
	}
	auth := v1.Group("")
	auth.Use(middleware.AuthRequired())
	{
		project := auth.Group("project")
		{
			project.GET("", controllers.ListProjest)
			project.POST("", controllers.CreateProject)
		}
		file := auth.Group("file")
		{
			file.POST("", controllers.Upload)
		}
		projectRelease := auth.Group("projectRelease")
		{
			projectRelease.POST("", controllers.CreateProjectRelease)
		}
		resources := v1.Group("resources")
		{
			resources.POST("", controllers.ListProjest)
		}
		task := v1.Group("task")
		{
			task.POST("", controllers.CreateTask)
		}
	}
	return r
}
func InitRouter() *gin.Engine {
	return InitMasterRouter()
}
