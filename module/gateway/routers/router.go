package routers

import (
	"github.com/wenchangshou2/vd-node-manage/module/gateway/g"
	middleware2 "github.com/wenchangshou2/vd-node-manage/module/gateway/middleware"
	controllers2 "github.com/wenchangshou2/vd-node-manage/module/gateway/routers/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// InitCORS 初始化跨域配置
func InitCORS(router *gin.Engine) {
	if g.Config().Cors.AllowOrigins[0] != "UNSET" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     g.Config().Cors.AllowOrigins,
			AllowMethods:     g.Config().Cors.AllowMethods,
			AllowHeaders:     g.Config().Cors.AllowHeaders,
			AllowCredentials: g.Config().Cors.AllowCredentials,
			ExposeHeaders:    g.Config().Cors.ExposeHeaders,
		}))
		return
	}

}
func InitMasterRouter() *gin.Engine {
	r := gin.Default()
	// InitCORS(r)
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
	r.StaticFS("/upload", http.Dir("./upload"))
	v1 := r.Group("/api/v1")
	v1.Use(middleware2.Session(g.Config().System.SessionSecret))
	v1.Use(middleware2.CurrentUser())
	{
		site := v1.Group("site")
		{
			site.GET("ping", controllers2.Ping)
		}
		user := v1.Group("user")
		{
			user.POST("session", controllers2.UserLogin)
			user.POST("",
				controllers2.UserRegister)
		}
		file := v1.Group("file")
		{
			file.GET(":id", controllers2.DownloadFile)
		}
		system := v1.Group("system")
		{
			system.GET("extranet", controllers2.GetExtranet)
		}
		computer := v1.Group("computer")
		{
			computer.PUT("", controllers2.UpdateComputer)
			computer.GET(":id/register", controllers2.GetComputerRegisterStatus)
			computer.GET(":id/details", controllers2.GetComputerDetails)
			computer.PUT(":id/name", controllers2.UpdateComputerName)
			computer.GET(":id/project", controllers2.ListComputerProject)
			computer.GET(":id/resource", controllers2.ListComputerResource)
			computer.GET(":id/task", controllers2.GetComputerTask)

			computer.DELETE(":id/resource/:resource_id", controllers2.DeleteComputerResource)
			computer.POST(":id/resource/:resource_id", controllers2.AddComputerResource)
			computer.GET(":id/projectRelease", controllers2.ListComputerProjectRelease)
			computer.POST(":id/projectRelease/:project_release_id", controllers2.AddComputerProjectRelease)
			computer.GET(":id/projectRelease/:project_release_id", controllers2.GetComputerProjectRelease)
			computer.DELETE(":id/projectRelease/:project_release_id", controllers2.DeleteComputerProjectRelease)
			computer.GET("", controllers2.ListComputer)
			computer.POST(":id/layout", controllers2.OpenMultiScreen)
			computer.POST(":id/:projectID/dir", controllers2.GetComputerProjectDir)
			computer.POST(":id/customLayout", controllers2.CreateCustomLayout)
			computer.GET(":id/customLayout", controllers2.GetComputerCustomLayout)
			computer.GET(":id/exhibition", controllers2.GetComputerExhibition)
		}
		exhibitionCategory := v1.Group("category")
		{
			exhibitionCategory.POST("", controllers2.CreateExhibitionCategory)
		}

		projectRelease := v1.Group("projectRelease")
		{
			projectRelease.POST(":id/computer/:computer_id", controllers2.AddComputerProjectRelease)
			projectRelease.GET(":id", controllers2.GetProjectRelease)
			projectRelease.POST(":id/publish", controllers2.PublishProject)
			projectRelease.GET(":id/file")
		}
		resources := v1.Group("resource")
		{
			resources.DELETE(":id", controllers2.DeleteResource)
			resources.POST(":id/publish", controllers2.PublishResource)
			resources.POST("list", controllers2.ListResource)
			resources.POST("", controllers2.CreateResource)
			resources.GET(":id/file", controllers2.DownloadResourceFile)
			resources.DELETE("id/file", controllers2.DeleteResourceFile)
		}
	}
	auth := v1.Group("")
	auth.Use(middleware2.AuthRequired())
	{
		user := v1.Group("user")
		{
			user.GET("currentUser", controllers2.GetCurrentUser)
		}
		computer := v1.Group("computer")
		{
			computer.GET("cross", controllers2.GetCrossResources)
			computer.POST(":id/exhibition", controllers2.OpenComputerExhibition)
			computer.POST(":id/report", controllers2.ReportComputerInfo)
			computer.GET(":id/heartbeat", controllers2.Heartbeat)
		}
		project := auth.Group("project")
		{
			project.GET("", controllers2.ListProject)
			project.GET(":id", controllers2.GetProjectReleaseList)
			project.POST("", controllers2.CreateProject)
		}
		file := auth.Group("file")
		{
			file.POST("", controllers2.Upload)
		}
		projectRelease := auth.Group("projectRelease")
		{
			projectRelease.POST("", controllers2.CreateProjectRelease)
			projectRelease.DELETE(":id", controllers2.DeleteProjectRelease)
		}

		task := v1.Group("task")
		{
			task.POST("list", controllers2.ListTask)
			task.POST("project", controllers2.CreateProjectTask)
			task.DELETE("project", controllers2.DeleteProjectTask)
			task.POST("resource", controllers2.CreateResourceTask)
			task.POST("", controllers2.UpdateTask)
			task.POST("taskItem", controllers2.UpdateTaskItem)
		}
		system := v1.Group("system")
		{
			system.GET("exportProjectRecord", controllers2.ExportProjectRecord)
		}

		exhibition := v1.Group("exhibition")
		{
			exhibition.POST("", controllers2.CreateComputerExhibition)
			exhibition.PUT("", controllers2.UpdateExhbition)
			exhibition.GET(":id", controllers2.GetExhibition)
			exhibition.DELETE(":id", controllers2.DeleteExhibition)
		}
		module := v1.Group("module")
		{
			module.GET("", controllers2.ListModule)
			module.DELETE(":id", controllers2.DeleteModule)
			module.POST("", controllers2.CreateModule)
		}
	}
	return r
}
func InitRouter() *gin.Engine {
	return InitMasterRouter()
}
