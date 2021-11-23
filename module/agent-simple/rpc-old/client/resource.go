package rpcClient

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/rpc-old/client/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ResourceServer struct {
	address string
}

func (server *ResourceServer) GetResourceDetailedinfo(resourceId string) (*pb.GetResourceDetailedInfoResponse, error) {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewResourceManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.GetResourceDetailedInfo(ctx, &wrapperspb.StringValue{
		Value: resourceId,
	})
	return response, err
}
func InitResourceServer(address string) *ResourceServer {
	return &ResourceServer{
		address: address,
	}
}
