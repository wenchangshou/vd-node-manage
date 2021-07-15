package task

import (
	"encoding/json"
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

// AddTaskItemService 添加下载任务项
type AddTaskItemService struct {
	Computers []uint `json:"computers" binding:"required"`
	Options   string `json:"options"`
	Action    uint   `json:"action" `
}

type InstallProjectOptions struct {
	Type             string `json:"type"`
	ProjectId        uint   `json:"project_id"`
	ProjectReleaseId uint   `json:"project_release_id"`
}
type InstallResourceOptions struct {
	ResourceId uint `json:"resource_id"`
}

type DeleteprojectOptions struct {
	Computers []uint `json:"computers" binding:"required"`
	Options   string `json:"options"`
	ID        uint   `json:"id"`
}

func (service AddTaskItemService) buildOptions() string {
	return ""

}

//installProject 安装项目
func (service *AddTaskItemService) installProject() serializer.Response {
	options := InstallProjectOptions{}
	err := json.Unmarshal([]byte(service.Options), &options)
	if err != nil {
		return serializer.Err(serializer.CodeJsonUnMarkshalErr, "解决安装参数失败", err)
	}
	pr, err := model.GetProjectReleaseByID(options.ProjectReleaseId)
	fmt.Println("pr", pr)
	if err != nil || pr.ID == 0 {
		return serializer.Err(serializer.CodeDBError, "获取项目版本失败", err)
	}
	// 全新安装
	if options.Type == "install" {
		for _, computer := range service.Computers {
			cp, err := model.GetComputerProjectById(uint(computer), options.ProjectId)
			// 如果当前的应用已经安装到远程的主机上面，需要先添加一个删除，之后再添加一个新的下载任务
			if err == nil && cp.ID > 0 {
				options := make(map[string]interface{})
				options["project_id"] = cp.ProjectId
				options["project_release_id"] = cp.ProjectReleaseId
				str, _ := json.Marshal(options)
				task := model.NewTask(int(computer), model.DeleteProject, string(str), 0)
				id, err := task.Add()
				if err != nil {
					return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
				}
				task = model.NewTask(int(computer), model.InstallProjectAction, service.Options, int(id))
				_, err = task.Add()
				if err != nil {
					return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
				}
			} else {

				task := model.NewTask(int(computer), model.InstallProjectAction, service.Options, 0)
				if id, err := task.Add(); err != nil && id == 0 {
					return serializer.Err(serializer.CodeDBError, "添加数据库任务失败", err)
				}
			}
		}
		return serializer.Response{
			Code: 0,
			Data: "OK",
		}
	}
	return serializer.Err(serializer.CodeParamErr, "未支持操作", nil)
}

// installResource 添加服务
func (service *AddTaskItemService) installResource() serializer.Response {
	options := InstallResourceOptions{}
	err := json.Unmarshal([]byte(service.Options), &options)
	if err != nil {
		return serializer.Err(serializer.CodeJsonUnMarkshalErr, "解决安装参数失败", err)
	}
	ids, err := model.AddTasks(service.Computers, service.Action, service.Options, false)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "添加任务项失败", err)
	}
	return serializer.Response{
		Data: ids,
	}

}

// Add 添加新的任务
func (service *AddTaskItemService) Add() serializer.Response {
	switch service.Action {
	case model.InstallResourceAction:
		return service.installResource()
	case model.InstallProjectAction:
		return service.installProject()
	default:
		return serializer.Err(serializer.CodeParamErr, "未支持的action", nil)
	}
}
