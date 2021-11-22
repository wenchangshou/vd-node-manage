package rpcServer

import (
	"context"

	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/rpc/server/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type SystemServer struct {
	pb.UnimplementedSystemManagementServer
}

func (s *SystemServer) Ping(ctx context.Context, request *pb.Empty) (wrapperspb.BoolValue, error) {
	// Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	return  wrapperspb.BoolValue{Value: true},nil
}
