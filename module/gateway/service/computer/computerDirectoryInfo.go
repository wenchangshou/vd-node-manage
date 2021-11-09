package computer

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server/pb"
)

type ProjectDirectoryService struct {
	ComputerID int
	ProjectID  int
	Dir        string `json:"dir"`
}
type GetDirectorForm struct {
	Action string `json:"action"`
	Dir    string `json:"Dir"`
}

func (service *ProjectDirectoryService) Get() serializer.Response {
	computer, err := model.GetComputerById(service.ComputerID)
	if err != nil || computer.ID == "" {
		return serializer.DBErr("获取计算机对象失败", err)
	}
	// project, err := model.GetComputerProjectByID(int(computer.ID), uint(service.ProjectID))
	// if err != nil {
	// 	return serializer.DBErr("获取项目失败", err)
	// }
	// projectRelease, err := model.GetComputerProjectByProjectIDAndProjectReleaseID(project.ID, project.ProjectReleaseID)
	// if err != nil || projectRelease == nil || len(projectRelease) == 0 {
	// 	return serializer.DBErr("获取项目发行版本失败", err)
	// }

	rpcServer.G_pubsubSerice.Publish(context.Background(), &pb.PublishChannel{Topic: computer.Ip, Id: "123456", Body: "{}", Action: "test"})
	return serializer.Response{}

}
