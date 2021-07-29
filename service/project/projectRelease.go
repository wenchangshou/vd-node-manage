package project

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/hashid"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
	"github.com/wenchangshou2/vd-node-manage/service/task"
)

type ProjectReleaseCreateService struct {
	Tag       string `form:"tag" json:"tag" binding:"required"`
	Content   string `form:"content" json:"content" binding:"required"`
	Mode      string `form:"mode" json:"mode" binding:"required"`
	Depend    uint   `form:"depend" json:"depend"`
	Arguments string `form:"arguments" json:"arguments"`
	Control   string `form:"control" json:"control"`
	FileId    uint   `form:"file_id" json:"file_id" binding:"required"`
	ProjectId uint   `form:"project_id" json:"project_id" binding:"required"`
}
type GetProjectReleaseService struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

func (service *GetProjectReleaseService) Get() serializer.Response {
	projectRelease, err := model.GetProjectReleaseByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNoFindProjectRelease, "没有找到相应的项目发行记录", err)
	}
	return serializer.Response{
		Data: projectRelease,
	}
}

func (service *ProjectReleaseCreateService) Create(c *gin.Context, user *model.User) serializer.Response {

	file, err := model.GetFileByUidAndId(int(service.FileId), user.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNoFindFileErr, "没有找到文件", err)
	}
	release := model.ProjectRelease{
		Tag:       service.Tag,
		Content:   service.Content,
		Mode:      service.Mode,
		Depend:    service.Depend,
		Arguments: service.Arguments,
		Control:   service.Control,
		FileID:    file.ID,
		ProjectID: service.ProjectId,
	}
	id, err := release.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建项目失败", err)
	}
	return serializer.Response{
		Data: hashid.HashID(id, hashid.ProjectID),
	}
}

type DeleteProejctReleaseService struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

func (service *DeleteProejctReleaseService) Delete() serializer.Response {
	projectRelease, err := model.GetProjectReleaseByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "没有找到对应的发行记录", err)
	}
	computerProject, err := model.GetComputerProjectByProjectIDAndProjectReleaseID(projectRelease.ProjectID, projectRelease.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机资源失败", err)
	}
	if len(computerProject) > 0 {
		ids := make([]uint, 0)
		for _, cp := range computerProject {
			ids = append(ids, cp.ComputerId)
		}
		taskItem := task.ComputerProject{
			Computers:        ids,
			ProjectID:        projectRelease.ProjectID,
			ProjectReleaseID: projectRelease.ID,
			Operator:         projectRelease.Mode,
		}
		response := taskItem.Delete()
		if response.Code != 0 {
			return response
		}
	}
	filePath := path.Join("upload/", projectRelease.File.SourceName)
	os.RemoveAll(filePath)
	projectRelease.File.Delete()
	projectRelease.Delete()

	return serializer.Response{}
}

type PublishProjectReleaseService struct {
	ID uint `uri:"id" json:"id"`
}

func (service *PublishProjectReleaseService) Publish() serializer.Response {
	projectRelease, err := model.GetProjectReleaseByID(service.ID)
	clientsids := make([]uint, 0)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机发行版本失败", err)
	}
	computers, _ := model.ListComputer()
	for _, computer := range computers {
		clientsids = append(clientsids, computer.ID)
	}

	taskItem := task.ComputerProject{
		Computers:        clientsids,
		ProjectID:        projectRelease.ProjectID,
		ProjectReleaseID: projectRelease.ID,
		Operator:         projectRelease.Mode,
	}
	return taskItem.Add()
}
