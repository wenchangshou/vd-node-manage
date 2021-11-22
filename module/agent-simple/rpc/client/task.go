package rpcClient

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/rpc/client/pb"
	"time"

	"google.golang.org/grpc"
)

type TaskServer struct {
	address string
}

func (server *TaskServer) GetTaskById(id string) error {
	return nil
}

// GenerateTaskItem 生成指定格式的任务列表
func (server *TaskServer) GetComputerTasksByMac(mac string, status pb.TaskType) ([]*pb.Task, error) {

	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewTaskManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.GetTaskByComputerMac(ctx, &pb.GetTaskRequest{
		Mac:      mac,
		TaskType: status,
	})
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}
func (server *TaskServer) SetTaskStatus(taskId string, status uint) (bool, error) {

	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return false, err
	}
	defer conn.Close()
	client := pb.NewTaskManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := pb.SetTaskStatusRequest{
		Id:   taskId,
		Type: pb.TaskType(status),
	}
	response, err := client.SetTaskStatus(ctx, &request)
	return response.GetValue(), err
}
func (server *TaskServer) SetTaskItemStatus(taskId string, status uint, msg string) (bool, error) {

	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return false, err
	}
	defer conn.Close()
	client := pb.NewTaskManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := pb.SetTaskItemStatusRequest{
		Id:   taskId,
		Type: pb.TaskType(status),
		Msg:  msg,
	}
	response, err := client.SetTaskItemStatus(ctx, &request)
	return response.GetValue(), err
}

// InitTaskServer 初始化任务服务
func InitTaskServer(address string) *TaskServer {
	return &TaskServer{
		address: address,
	}
}
