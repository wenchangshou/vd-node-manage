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

func (t TaskRpcService) SetTaskItemStatus(ids []uint, i int) error {

	panic("implement me")
}

func (t TaskRpcService) SetTaskStatus(ids []uint, status model2.EventStatus) error {
	client := &g.SingleConnRpcClient{
		RpcServer: fmt.Sprintf(g.Config().Server.RpcAddress),
		Timeout:   time.Second,
	}
	req := model2.DeviceSetStatusRequest{
		ID:     ids,
		Status: status,
	}
	response := model2.SimpleRpcResponse{}
	err := client.Call("Task.SetResourceDistributionStatus", req, &response)
	if err != nil {
		return err
	}
	panic("implement me")
}

// GetTasks 获取任务栏
func (t TaskRpcService) GetTasks() ([]dto.Task, error) {
	//rpcClient := &g.SingleConnRpcClient{
	//	RpcServer: fmt.Sprintf(g.Config().Server.RpcAddress),
	//	Timeout:   time.Second,
	//}
	//req := model2.QueryDeviceResourceDistributionRequest{DeviceID: t.ID}
	//reply := model2.QueryDeviceResourceDistributionResponse{}
	//err := rpcClient.Call("Task.QueryDeviceResourceDistribution", &req, &reply)
	//return reply.Tasks, err
	return nil, nil
}

// NewTaskRpcService 新的任务rpc服务
func NewTaskRpcService(id uint) *TaskRpcService {
	return &TaskRpcService{
		id,
	}
}
