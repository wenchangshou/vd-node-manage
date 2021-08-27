package rpcServer

import (
	"context"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/rpc/server/pb"
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
func (s *TaskServer) ConvertTaskJsonToProtoBuf(task *model.Task, item []model.TaskItem) *pb.Task {
	result := &pb.Task{}
	result.Name = task.Name
	result.Id = int32(task.ID)
	result.TaskItem = make([]*pb.TaskItem, 0)
	for _, _item := range item {
		taskItem := &pb.TaskItem{}
		taskItem.Action = pb.TaskOperatorType(_item.Action)
		taskItem.Depend = int32(_item.Depend)
		taskItem.Id = int32(_item.ID)
		taskItem.Options = _item.Options
		taskItem.Schedule = int32(_item.Schedule)
		taskItem.Status = int32(_item.Status)
		result.TaskItem = append(result.TaskItem, taskItem)
	}
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
		subItem, _ := model.GetTaskItemById(int(v.ID))
		item := s.ConvertTaskJsonToProtoBuf(&v, subItem)
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

func (s TaskServer) SetTaskItemStatus(ctx context.Context, request *pb.SetTaskItemStatusRequest) (*wrapperspb.BoolValue, error) {
	err := model.SetTaskItemStatus(uint(request.GetId()), uint(request.GetType()), request.Msg)
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}
