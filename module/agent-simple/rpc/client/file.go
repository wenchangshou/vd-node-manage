package rpcClient

import (
	"context"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/rpc/client/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type FileServer struct {
	address string
}

func (server *FileServer) GetFileInfoByProjectReleaseID(id string) (*pb.GetFileInfoByProjectReleaseIDResponse, error) {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewFileManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.GetFileInfoByProjectReleaseID(ctx, &wrapperspb.StringValue{
		Value: id,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("获取的文件信息:%v\n", response)
	return response, nil
}

func InitFileServer(address string) *FileServer {
	return &FileServer{
		address: address,
	}

}
