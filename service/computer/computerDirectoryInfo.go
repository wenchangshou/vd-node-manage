package computer

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type ComputerProjectDirectoryService struct {
	ComputerID int
	ProjectID  int
	Dir        string `json:"dir"`
}
type GetDirectorForm struct {
	Action string `json:"action"`
	Dir    string `json:"Dir"`
}

func (service *ComputerProjectDirectoryService) Get() serializer.Response {
	computer, err := model.GetComputerById(service.ComputerID)
	if err != nil || computer.ID == 0 {
		return serializer.DBErr("获取计算机对象失败", err)
	}
	project, err := model.GetComputerProjectByID(int(computer.ID), uint(service.ProjectID))
	if err != nil {
		return serializer.DBErr("获取项目失败", err)
	}
	projectRelease, err := model.GetComputerProjectByProjectIDAndProjectReleaseID(project.ID, project.ProjectReleaseID)
	if err != nil || projectRelease == nil || len(projectRelease) == 0 {
		return serializer.DBErr("获取项目发行版本失败", err)
	}

}
