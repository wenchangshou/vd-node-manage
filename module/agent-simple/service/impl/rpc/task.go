package rpc

import (
	"fmt"
	"time"

	model2 "github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
)

type TaskRpcService struct {
	ID uint `json:"id"`
}

func (t TaskRpcService) SetTaskItemStatus(_ []uint, _ int) error {

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

// NewTaskRpcService 新的任务rpc服务
func NewTaskRpcService(id uint) *TaskRpcService {
	return &TaskRpcService{
		id,
	}
}
