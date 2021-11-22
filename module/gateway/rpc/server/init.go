package rpcServer

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server/pb"
	"net"

	"google.golang.org/grpc"
)

func InitRpc(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logging.GLogger.Error(fmt.Sprintf("failed to listen:%v", err))
		return err
	}
	s := grpc.NewServer()
	pb.RegisterTaskManagementServer(s, &TaskServer{})
	pb.RegisterSystemManagementServer(s, &SystemServer{})
	pb.RegisterFileManagementServer(s, &FileServer{})
	pb.RegisterComputerManagementServer(s, &ComputerServer{})
	pb.RegisterResourceManagementServer(s, &ResourceServer{})
	pb.RegisterPubsubServiceServer(s, NewPubsService())

	if err := s.Serve(lis); err != nil {
		logging.GLogger.Error(fmt.Sprintf("grpc new server error:%v", err))
		return err
	}
	return nil
}