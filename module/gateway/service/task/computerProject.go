package task

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

// ComputerProject 计算机项目
type ComputerProject struct {
	Computers        []string `json:"computers" binding:"required"`
	Operator         string   `json:"operator" `
	ProjectID        string   `json:"project_id"`
	ProjectReleaseID string   `json:"project_release_id" binding:"required"`
}
type DeleteProjectOption struct {
	Id   uint32
	File *model.File
}

//install 安装一个新的资源
func (service ComputerProject) install() serializer.Response {
	option := make(map[string]interface{})
	projectRelease, err := model.GetProjectReleaseByID(service.ProjectReleaseID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取对应的发行版本失败", err)
	}
	option["id"] = projectRelease.ID
	option["uri"] = "upload/" + projectRelease.File.SourceName
	option["name"] = projectRelease.File.Name
	option["source"] = projectRelease.File.SourceName
	// 根据项目id，获取所有的发行id
	projectReleaseIds := make([]string, 0)
	projectReleaseList, err := model.GetProjectReleaseListByProjectID(service.ProjectID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取项目对应的发行版本列表失败", err)
	}
	for _, pr := range projectReleaseList {
		projectReleaseIds = append(projectReleaseIds, pr.ID)
	}
	for _, computer := range service.Computers {
		depend := ""
		pr, err := model.CheckComputerProjectRelease(computer, projectReleaseIds)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "检查是否存在项目错误", err)
		}
		task, err := model.AddTask(fmt.Sprintf("添加%s项目", service.ProjectReleaseID), computer)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
		}
		if len(pr) > 0 {
			_option := make(map[string]interface{})
			_option["id"] = service.ProjectReleaseID
			_option["file"] = pr[0].File
			subTask, err := model.AddTaskItem(task.ID, model.DeleteProject, _option, false, "")
			if err != nil {
				return serializer.Err(serializer.CodeDBError, "添加子任务失败", err)
			}
			depend = subTask.ID
		}
		model.AddTaskItem(task.ID, model.InstallProjectAction, option, false, depend)
	}

	return serializer.Response{
		Code: 0,
		Data: "OK",
	}
}
func (service ComputerProject) Add() serializer.Response {
	projectRelease, err := model.GetProjectReleaseByID(service.ProjectReleaseID)
	if err != nil || projectRelease.ID == "" {
		return serializer.Err(serializer.CodeJsonUnMarkshalErr, "获取项目发行版本失败", err)
	}
	switch service.Operator {
	case "install":
		return service.install()
	default:
		return serializer.Err(serializer.CodeNotSupportOperator, "未支持的操作", err)
	}

}
func (service ComputerProject) Delete() serializer.Response {
	// for _, computer := range service.Computers {
	// 	service, err := model.GetComputerProjectByID(computer, service.ProjectID)
	// 	if err != nil || service.ID == "" {
	// 		return serializer.Err(serializer.CodeNotFindComputerProject, "没有找到对应的计算机项目", err)
	// 	}
	// 	projectRelease, err := model.GetProjectReleaseByID(service.ProjectReleaseID)
	// 	if err != nil {
	// 		return serializer.Err(serializer.CodeNotFindProjectRelease, "没有找到对应的项目发布版本", err)
	// 	}
	// 	options := make(map[string]interface{})
	// 	options["ID"] = projectRelease.ID
	// 	options["File"] = projectRelease.File

	// 	task, err := model.AddTask(fmt.Sprintf("删除%s项目", projectRelease.Project.Name), computer)
	// 	if err != nil {
	// 		return serializer.Err(serializer.CodeDBError, "创建任务失败", err)
	// 	}
	// 	model.AddTaskItem(task.ID, model.DeleteProject, options, false, "")
	// }
	return serializer.Response{
		Code: 0,
		Data: "OK",
	}
}
