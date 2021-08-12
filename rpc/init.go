package rpc

import (
	"fmt"
	"net"

	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"github.com/wenchangshou2/vd-node-manage/rpc/pb"
	"google.golang.org/grpc"
)

func InitRpc(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logging.G_Logger.Error(fmt.Sprintf("failed to listen:%v", err))
		return err
	}
	s := grpc.NewServer()
	pb.RegisterTaskManagementServer(s, &TaskServer{})
	pb.RegisterSystemManagementServer(s, &SystemServer{})
	pb.RegisterFileManagementServer(s, &FileServer{})
	pb.RegisterComputerManagementServer(s, &ComputerServer{})
	pb.RegisterResourceManagementServer(s, &ResourceServer{})
	pb.RegisterPubsubServiceServer(s, NewPubsubService())

	if err := s.Serve(lis); err != nil {
		logging.G_Logger.Error(fmt.Sprintf("grpc new server error:%v", err))
		return err
	}
	return nil
}
