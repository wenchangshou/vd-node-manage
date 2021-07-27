package routers

import (
	"net/http"

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
	r.StaticFS("/upload", http.Dir("./upload"))
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
		file := v1.Group("file")
		{
			file.GET(":id", controllers.DownloadFile)
		}
		system := v1.Group("system")
		{
			system.GET("extranet", controllers.GetExtranet)
		}
		client := v1.Group("computer")
		{
			client.PUT("", controllers.UpdateComputer)
			client.GET(":id/details", controllers.GetComputerDetails)
			client.PUT(":id/name", controllers.UpdateComputerName)
			client.GET(":id/project", controllers.ListComputerProject)
			client.GET("", controllers.ListComputer)
		}

		projectRelease := v1.Group("projectRelease")
		{
			projectRelease.GET(":id", controllers.GetProjectRelease)
			projectRelease.POST(":id/publish", controllers.PublishProject)
		}
	}
	auth := v1.Group("")
	auth.Use(middleware.AuthRequired())
	{
		user := v1.Group("user")
		{
			user.GET("currentUser", controllers.GetCurrentUser)
		}
		project := auth.Group("project")
		{
			project.GET("", controllers.ListProjest)
			project.GET(":id", controllers.GetProjectReleaseList)
			project.POST("", controllers.CreateProject)
		}
		file := auth.Group("file")
		{
			file.POST("", controllers.Upload)
		}
		projectRelease := auth.Group("projectRelease")
		{
			projectRelease.POST("", controllers.CreateProjectRelease)
			projectRelease.DELETE(":id", controllers.DeleteProjectRelease)
		}
		resources := v1.Group("resource")
		{
			resources.GET("", controllers.ListResource)
			resources.POST("", controllers.CreateResource)
		}
		task := v1.Group("task")
		{
			task.POST("project", controllers.CreateProjectTask)
			task.DELETE("project", controllers.DeleteProjectTask)
			task.POST("resource", controllers.CreateResourceTask)
			task.GET("", controllers.ListTask)
		}
	}
	return r
}
func InitRouter() *gin.Engine {
	return InitMasterRouter()
}
