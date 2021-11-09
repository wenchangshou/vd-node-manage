package rpcServer

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server/pb"
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
	resource, err := model.GetResourceById(request.ResourceID)
	if err != nil {
		return nil, err
	}
	err = computer.AppendNewResource(*resource)
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}
func (server *ComputerServer) AddComputerProject(ctx context.Context, request *pb.SetComputerProjectRequest) (*wrapperspb.BoolValue, error) {
	// computer, err := model.GetComputerByMac(request.Mac)
	// if err != nil {
	// 	return nil, err
	// }
	// newproject := &model.ComputerProject{
	// 	ComputerId:       computer.ID,
	// 	ProjectID:        request.ProjectID,
	// 	ProjectReleaseID: request.ProjectReleaseID,
	// }
	// id, err := newproject.Create()
	// return &wrapperspb.BoolValue{
	// 	Value: err == nil && id != "",
	// }, err
	return &wrapperspb.BoolValue{
		Value: true,
	}, nil
}

//GetComputerProject 获取计算机项目
func (server *ComputerServer) GetComputerProject(ctx context.Context, request *pb.GetComputerProjectRequest) (response *pb.GetComputerProjectResponse, err error) {
	// return &wrapperspb.BoolValue{
	// 	Value: true,
	// }, nil
	// computerProject, err := model.GetComputerProjectByID(request.ComputerId, request.Id)
	// if err != nil {
	// 	return nil, err
	// }
	// projectRelease, err := model.GetProjectReleaseByID(computerProject.ProjectReleaseID)
	// if err != nil {
	// 	return nil, err
	// }
	// response = &pb.GetComputerProjectResponse{
	// 	Id: request.Id,
	// 	File: &pb.File{
	// 		ID:         projectRelease.FileID,
	// 		Name:       projectRelease.File.Name,
	// 		SourceName: projectRelease.File.SourceName,
	// 		Size:       int64(projectRelease.File.Size),
	// 		Url:        "upload/" + projectRelease.File.SourceName,
	// 		Uuid:       projectRelease.File.Uuid,
	// 	},
	// }
	return
}
func (server *ComputerServer) DeleteComputerProject(ctx context.Context, id *wrapperspb.UInt32Value) (status *wrapperspb.BoolValue, err error) {
	// err = model.DeleteComputerProjectByID(int(id.GetValue()))
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}
func (server *ComputerServer) DeleteComputerResource(ctx context.Context, id *wrapperspb.UInt32Value) (status *wrapperspb.BoolValue, err error) {
	// err = model.DeleteComputerResourceById(int(id.GetValue()))
	// return &wrapperspb.BoolValue{
	// 	Value: err == nil,
	// }, err
	return &wrapperspb.BoolValue{
		Value: true,
	}, nil
}

func (server *ComputerServer) GetComputerByMac(ctx context.Context, mac wrapperspb.StringValue) (id *wrapperspb.StringValue, err error) {
	computer, err := model.GetComputerByMac(mac.GetValue())
	return &wrapperspb.StringValue{
		Value: computer.ID,
	}, err
}
