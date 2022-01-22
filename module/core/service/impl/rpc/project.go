package rpc

import (
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
)

type ProjectRpcService struct {
	ID     uint
	Client *g.SingleConnRpcClient
}

func (service ProjectRpcService) QueryProject(id uint) (project *model.ProjectInfo, err error) {
	req := model.NormalIdRequest{ID: id}
	reply := model.ProjectQueryResponse{}
	if err = service.Client.Call("Project.Query", &req, &reply); err != nil {
		return nil, err
	}
	return &reply.Project, nil
}
func NewProjectRpcService(id uint, client *g.SingleConnRpcClient) *ProjectRpcService {
	return &ProjectRpcService{
		ID:     id,
		Client: client,
	}
}
