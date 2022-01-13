package rpc

import (
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
)

type ResourceRpcService struct {
	ID     uint
	Client *g.SingleConnRpcClient
}

func (service ResourceRpcService) QueryResource(id uint) (*model.ResourceInfo, error) {
	var (
		err error
	)
	req := model.ResourceQueryRequest{ID: id}
	reply := model.ResourceQueryResponse{}
	if err = service.Client.Call("Resource.Query", &req, &reply); err != nil {
		return nil, err
	}

	return &reply.Resource, nil

}
func NewResourceRpcService(id uint, client *g.SingleConnRpcClient) *ResourceRpcService {
	return &ResourceRpcService{
		ID:     id,
		Client: client,
	}
}
