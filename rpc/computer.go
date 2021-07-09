package rpc

import (
	"context"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ComputerServer struct {
	pb.UnimplementedComputerManagementServer
}

func (server *ComputerServer) AddComputerResource(ctx context.Context, request *pb.SetComputerResourceRequest) (*wrapperspb.BoolValue, error) {
	computer, err := model.GetComputerByMac(request.Mac)
	if err != nil {
		return nil, err
	}
	resource := &model.ComputerResource{
		ComputerId: computer.ID,
		ResourceId: uint(request.ResourceID),
	}
	id, err := resource.Create()
	return &wrapperspb.BoolValue{
		Value: err == nil && id > 0,
	}, err
}
func (server *ComputerServer) AddComputerProject(ctx context.Context, request *pb.SetComputerProjectRequest) (*wrapperspb.BoolValue, error) {
	computer, err := model.GetComputerByMac(request.Mac)
	if err != nil {
		return nil, err
	}
	newproject := &model.ComputerProject{
		ComputerId:       computer.ID,
		ProjectId:        uint(request.ProjectID),
		ProjectReleaseId: uint(request.ProjectReleaseID),
	}
	id, err := newproject.Create()
	return &wrapperspb.BoolValue{
		Value: err == nil && id > 0,
	}, err
}
