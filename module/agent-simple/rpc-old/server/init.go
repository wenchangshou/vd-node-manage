package rpcServer

import (
	"net"

	"google.golang.org/grpc"
)

func InitRpc(connStr string) error {
	lis, err := net.Listen("tcp", connStr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterSystemManagementServer(s, &SystemServer{})
}
