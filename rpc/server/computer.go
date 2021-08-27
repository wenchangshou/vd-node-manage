package rpcServer

import (
	"context"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/rpc/server/pb"
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
		ComputerID: computer.ID,
		ResourceID: uint(request.ResourceID),
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
		ProjectID:        uint(request.ProjectID),
		ProjectReleaseID: uint(request.ProjectReleaseID),
	}
	id, err := newproject.Create()
	return &wrapperspb.BoolValue{
		Value: err == nil && id > 0,
	}, err
}

//GetComputerProject 获取计算机项目
func (server *ComputerServer) GetComputerProject(ctx context.Context, request *pb.GetComputerProjectRequest) (response *pb.GetComputerProjectResponse, err error) {
	computerProject, err := model.GetComputerProjectByID(int(request.ComputerId), uint(request.Id))
	if err != nil {
		return nil, err
	}
	projectRelease, err := model.GetProjectReleaseByID(computerProject.ProjectReleaseID)
	if err != nil {
		return nil, err
	}
	response = &pb.GetComputerProjectResponse{
		Id: request.Id,
		File: &pb.File{
			ID:         int32(projectRelease.FileID),
			Name:       projectRelease.File.Name,
			SourceName: projectRelease.File.SourceName,
			Size:       int64(projectRelease.File.Size),
			Url:        "upload/" + projectRelease.File.SourceName,
			Uuid:       projectRelease.File.Uuid,
		},
	}
	return
}
func (server *ComputerServer) DeleteComputerProject(ctx context.Context, id *wrapperspb.UInt32Value) (status *wrapperspb.BoolValue, err error) {
	err = model.DeleteComputerProjectByID(int(id.GetValue()))
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}
func (server *ComputerServer) DeleteComputerResource(ctx context.Context, id *wrapperspb.UInt32Value) (status *wrapperspb.BoolValue, err error) {
	err = model.DeleteComputerResourceById(int(id.GetValue()))
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}

func (server *ComputerServer) GetComputerByMac(ctx context.Context, mac wrapperspb.StringValue) (id *wrapperspb.UInt32Value, err error) {
	computer, err := model.GetComputerByMac(mac.GetValue())
	return &wrapperspb.UInt32Value{
		Value: uint32(computer.ID),
	}, err
}
