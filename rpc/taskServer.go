package rpc

import (
	"context"

	"github.com/wenchangshou2/vd-node-manage/model"
	pb "github.com/wenchangshou2/vd-node-manage/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type TaskServer struct {
	pb.UnimplementedTaskManagementServer
}

func (s *TaskServer) getTaskTypeByNumber(t uint) pb.TaskOperatorType {
	if t == model.InstallProjectAction {
		return pb.TaskOperatorType_INSTALL_PROJECT
	} else if t == model.InstallResourceAction {
		return pb.TaskOperatorType_INSTALL_RESOURCE
	} else if t == model.DeleteResource {
		return pb.TaskOperatorType_DELETE_RESOURCE
	} else if t == model.DeleteProject {
		return pb.TaskOperatorType_DELETE_PROJECT
	}
	return pb.TaskOperatorType_UNKNOWN
}
func (s *TaskServer) ConvertTaskJsonToProtoBuf(task *model.Task) *pb.Task {
	result := &pb.Task{}
	result.Action = s.getTaskTypeByNumber(task.Action)
	result.Depend = int32(task.Depend)
	result.Options = task.Options
	result.Schedule = int32(task.Schedule)
	result.Status = int32(task.Status)
	result.ID = int32(task.ID)
	return result
}
func (s *TaskServer) GetTaskByComputerMac(ctx context.Context, request *pb.GetTaskRequest) (*pb.TasksResponse, error) {
	tasks := make([]model.Task, 0)
	response := &pb.TasksResponse{}
	items := make([]*pb.Task, 0)
	computer, err := model.GetComputerByMac(request.Mac)
	if err != nil {
		return response, err
	}
	if request.TaskType == pb.TaskType_ALL {
		tasks, err = model.GetTaskListByCid(int(computer.ID))
	} else {
		tasks, err = model.GetTaskListByCidFilterStatus(int(computer.ID), int(request.TaskType))
	}
	if err != nil {
		return response, err
	}
	for _, v := range tasks {
		item := s.ConvertTaskJsonToProtoBuf(&v)
		items = append(items, item)
	}
	response.Items = items
	return response, nil
}
func (s TaskServer) SetTaskStatus(ctx context.Context, request *pb.SetTaskStatusRequest) (*wrapperspb.BoolValue, error) {
	err := model.SetTaskStatus(uint(request.GetId()), uint(request.GetType()))
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}
