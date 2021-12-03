package rpc

import (
	"fmt"
	model2 "github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/dto"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"time"
)

type TaskRpcService struct {
	ID uint `json:"id"`
}

func (t TaskRpcService) SetTaskItemStatus(strings []string, i int) error {
	panic("implement me")
}

func (t TaskRpcService) SetTaskStatus(strings []string, i int) error {
	panic("implement me")
}

func (t TaskRpcService) GetTasks() ([]dto.Task, error) {
	rpcClient := &g.SingleConnRpcClient{
		RpcServer: fmt.Sprintf(g.Config().Server.RpcAddress),
		Timeout:   time.Second,
	}
	req := model2.QueryDeviceResourceDistributionRequest{DeviceID: t.ID}
	reply := model2.QueryDeviceResourceDistributionResponse{}
	err := rpcClient.Call("Task.QueryDeviceResourceDistribution", &req, &reply)
	return reply.Tasks, err
}
func NewTaskRpcService(id uint) *TaskRpcService {
	return &TaskRpcService{
		id,
	}
}
