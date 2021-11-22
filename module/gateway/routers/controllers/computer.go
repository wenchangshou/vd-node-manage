package controllers

import (
	"fmt"
	computer2 "github.com/wenchangshou2/vd-node-manage/module/gateway/service/computer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateComputer(c *gin.Context) {
	var service computer2.UpdateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func UpdateComputerName(c *gin.Context) {
	var service computer2.UpdateNameService
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	service.ID = c.Param("id")
	res := service.Update()
	c.JSON(200, res)
}
func ListComputer(c *gin.Context) {
	var service computer2.ListService
	res := service.List()
	c.JSON(200, res)
}

// ListComputerResource 获取指定计算机的资源列表
func ListComputerResource(c *gin.Context) {
	var service computer2.ListComputerResourceService
	if err := c.ShouldBindQuery(&service); err == nil {
		id := c.Param("id")
		service.ID = id
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
func ListComputerProject(c *gin.Context) {
	var service computer2.ComputerProjectListService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
func GetComputerRegisterStatus(c *gin.Context) {

	var service computer2.IDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.IsRegister()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetComputerDetails(c *gin.Context) {
	var service computer2.IDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func GetCrossResources(c *gin.Context) {
	var service computer2.ComputerProjectGetCrossResource
	res := service.Get()
	c.JSON(200, res)
}
func OpenComputerExhibition(c *gin.Context) {
	var service computer2.ComputerExhibitionOpenService
	if err := c.ShouldBindUri(service); err == nil {
	}
}
func GetComputerProjectDir(c *gin.Context) {
	computerID := c.Param("id")
	projectID := c.Param("projectID")
	fmt.Printf("computerID:%v,projectID:%v", computerID, projectID)
	var service computer2.ProjectDirectoryService
	if err := c.ShouldBindJSON(&service); err == nil {
		service.ComputerID, _ = strconv.Atoi(computerID)
		service.ProjectID, _ = strconv.Atoi(projectID)
		service.Get()
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func AddComputerResource(c *gin.Context) {
	var service computer2.ComputerResourceService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func DeleteComputerResource(c *gin.Context) {
	var service computer2.ComputerResourceService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func AddComputerProjectRelease(c *gin.Context) {
	var service computer2.ProjectReleaseService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteComputerProjectRelease 删除计算机项目
func DeleteComputerProjectRelease(c *gin.Context) {
	var service computer2.ProjectReleaseService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// GetComputerProjectRelease 获取计算机指定项目
func GetComputerProjectRelease(c *gin.Context) {
	var service computer2.ProjectReleaseService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListComputerProjectRelease 获取计算机项目列表
func ListComputerProjectRelease(c *gin.Context) {
	var service computer2.ProjectReleaseService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetComputerTask(c *gin.Context) {
	var service computer2.ListComputerTaskService
	if err := c.ShouldBind(&service); err == nil {
		service.ID = c.Param("id")
		res := service.GetComputerTask()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ReportComputerInfo 上报计算机服务信息
func ReportComputerInfo(c *gin.Context) {
	var service computer2.UpdateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func Heartbeat(c *gin.Context) {
	var service computer2.IDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Heartbeat()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}