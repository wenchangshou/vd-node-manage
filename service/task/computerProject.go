package task

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

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
func (computerProject ComputerProject) install() serializer.Response {
	// for _, computer := range computerProject.Computers {
	// 	depend := ""
	// 	projectRelease, err := model.GetProjectReleaseByIdAndProjectId(computerProject.ProjectID, computerProject.ProjectReleaseID)
	// 	if err != nil {
	// 		return serializer.Err(serializer.CodeNotFindProjectRelease, "没有找到项目", err)
	// 	}
	// 	options := make(map[string]interface{})
	// 	options["url"] = "upload/" + projectRelease.File.SourceName
	// 	options["project_id"] = computerProject.ProjectID
	// 	options["project_release_id"] = computerProject.ProjectReleaseID
	// 	options["File"] = projectRelease.File
	// 	cp, _ := model.GetComputerProjectByID(computer, computerProject.ProjectID)
	// 	task, err := model.AddTask(fmt.Sprintf("添加%s项目", projectRelease.Project.Name), computer)
	// 	if err != nil {
	// 		return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
	// 	}
	// 	if err == nil && cp.ID != "" {
	// 		_options := make(map[string]interface{})
	// 		_options["ID"] = cp.ID
	// 		_projectRelease, _ := model.GetProjectReleaseByID(cp.ProjectReleaseID)
	// 		_options["File"] = _projectRelease.File
	// 		task, err := model.AddTaskItem(task.ID, model.DeleteProject, _options, false, "")
	// 		if err != nil {
	// 			return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
	// 		}
	// 		depend = task.ID
	// 	}
	// 	model.AddTaskItem(task.ID, model.InstallProjectAction, options, false, depend)

	// }
	return serializer.Response{
		Code: 0,
		Data: "OK",
	}
}
func (computerProject ComputerProject) Add() serializer.Response {
	projectRelease, err := model.GetProjectReleaseByID(computerProject.ProjectReleaseID)
	if err != nil || projectRelease.ID == "" {
		return serializer.Err(serializer.CodeJsonUnMarkshalErr, "获取项目发行版本失败", err)
	}
	switch computerProject.Operator {
	case "install":
		return computerProject.install()
	default:
		return serializer.Err(serializer.CodeNotSupportOperator, "未支持的操作", err)
	}

}
func (computerProject ComputerProject) Delete() serializer.Response {
	// for _, computer := range computerProject.Computers {
	// 	computerProject, err := model.GetComputerProjectByID(computer, computerProject.ProjectID)
	// 	if err != nil || computerProject.ID == "" {
	// 		return serializer.Err(serializer.CodeNotFindComputerProject, "没有找到对应的计算机项目", err)
	// 	}
	// 	projectRelease, err := model.GetProjectReleaseByID(computerProject.ProjectReleaseID)
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
