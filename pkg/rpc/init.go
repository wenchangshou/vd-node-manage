package rpc

import (
	"fmt"
	"net"

	pb "github.com/wenchangshou2/vd-node-manage/pb"
	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
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

	if err := s.Serve(lis); err != nil {
		logging.G_Logger.Error(fmt.Sprintf("grpc new server error:%v", err))
		return err
	}
	return nil
}
