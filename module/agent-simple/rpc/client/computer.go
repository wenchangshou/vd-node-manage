package rpcClient

import (
	"context"
	"errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/rpc/client/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ComputerServer struct {
	address string
}

// 获取计算机列表
func (server *ComputerServer) GetClient() (*pb.ComputerManagementClient, error) {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewComputerManagementClient(conn)
	return &client, err
}

// AddComputerProject 添加计算机项目
func (server *ComputerServer) AddComputerProject(mac string, projectId string, projectReleaseId string) error {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewComputerManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.AddComputerProject(ctx, &pb.SetComputerProjectRequest{
		Mac:              mac,
		ProjectID:        projectId,
		ProjectReleaseID: projectReleaseId,
	})
	if err != nil {
		return err
	}
	if !response.Value {
		return errors.New("添加计算机项目失败")
	}
	return nil
}

// 删除计算机项目
func (server *ComputerServer) DeleteComputerProject(id uint32) error {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewComputerManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.DeleteComputerProject(ctx, &wrapperspb.UInt32Value{Value: id})
	return err
}
func (server *ComputerServer) DeleteComputerResource(id uint32) error {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewComputerManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.DeleteComputerResource(ctx, &wrapperspb.UInt32Value{Value: id})
	return err
}

// AddComputerResource 添加计算机资源
func (server *ComputerServer) AddComputerResource(mac string, resourceId string) error {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewComputerManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.AddComputerResource(ctx, &pb.SetComputerResourceRequest{
		Mac:        mac,
		ResourceID: resourceId,
	})
	if err != nil {
		return err
	}
	if !response.Value {
		return errors.New("添加计算机项目失败")
	}
	return nil
}

// GetComputerIdByMac 通过mac获取对应的计算机id
func (server *ComputerServer) GetComputerIdByMac(mac string) (uint, error) {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return 0, err
	}
	client := pb.NewComputerManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	id, err := client.GetComputerIdByMac(ctx, &wrapperspb.StringValue{
		Value: mac,
	})
	return uint(id.GetValue()), err
}

// GetComputerProject 获取计算机项目
func (server *ComputerServer) GetComputerProject(computerID, projectID string) (response *pb.GetComputerProjectResponse, err error) {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewComputerManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return client.GetComputerProject(ctx, &pb.GetComputerProjectRequest{
		Id:         projectID,
		ComputerId: computerID,
	})
}

func InitComputerServer(address string) *ComputerServer {
	return &ComputerServer{
		address: address,
	}
}
