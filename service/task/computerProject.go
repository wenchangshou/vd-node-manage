package task

import (
	"encoding/json"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ComputerProject struct {
	Computers        []uint `json:"computers" binding:"required"`
	Operator         string `json:"operator" `
	ProjectID        uint   `json:"project_id"`
	ProjectReleaseID uint   `json:"project_release_id" binding:"required"`
}
type DeleteProjectOption struct {
	Id   uint32
	File *model.File
}

func (computerProject ComputerProject) install() serializer.Response {
	for _, computer := range computerProject.Computers {
		projectRelease, err := model.GetProjectReleaseByIdAndProjectId(computerProject.ProjectReleaseID, computerProject.ProjectID)
		if err != nil {
			return serializer.Err(serializer.CodeNotFindProjectRelease, "没有找到项目", err)
		}
		options := make(map[string]interface{})
		options["url"] = "upload/" + projectRelease.File.SourceName
		options["project_id"] = computerProject.ProjectID
		options["project_release_id"] = computerProject.ProjectReleaseID
		options["File"] = projectRelease.File
		str, _ := json.Marshal(options)
		cp, err := model.GetComputerProjectById(computer, computerProject.ProjectID)
		if err == nil && cp.ID > 0 {
			task := model.NewTask(int(computer), model.DeleteProject, string(str), 0)
			id, err := task.Add()
			if err != nil {
				return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
			}
			task = model.NewTask(int(computer), model.InstallProjectAction, string(str), int(id))
			_, err = task.Add()
			if err != nil {
				return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
			}
		} else {
			task := model.NewTask(int(computer), model.InstallProjectAction, string(str), 0)
			_, err = task.Add()
			if err != nil {
				return serializer.Err(serializer.CodeDBError, "添加任务失败", err)
			}
		}
	}
	return serializer.Response{
		Code: 0,
		Data: "OK",
	}
}
func (computerProject ComputerProject) Add() serializer.Response {
	projectRelease, err := model.GetProjectReleaseByID(computerProject.ProjectReleaseID)
	if err != nil || projectRelease.ID == 0 {
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
	for _, computer := range computerProject.Computers {
		computerProject, err := model.GetComputerProjectById(computer, computerProject.ProjectID)
		if err != nil || computerProject.ID == 0 {
			return serializer.Err(serializer.CodeNotFindComputerProject, "没有找到对应的计算机项目", err)
		}
		projectRelease, err := model.GetProjectReleaseByID(computerProject.ProjectReleaseId)
		if err != nil {
			return serializer.Err(serializer.CodeNotFindProjectRelease, "没有找到对应的项目发布版本", err)
		}
		options := DeleteProjectOption{
			Id:   uint32(projectRelease.ID),
			File: &projectRelease.File,
		}
		jsonStr, _ := json.Marshal(options)
		task := model.NewTask(int(computer), model.DeleteProject, string(jsonStr), 0)
		_, err = task.Add()
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "创建任务失败", err)
		}

	}
	return serializer.Response{
		Code: 0,
		Data: "OK",
	}
}
