package rpcClient

import (
	"context"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/rpc-old/client/pb"
	"os"
	"time"

	"google.golang.org/grpc"
)

type SystemServer struct {
	address string
}

func (server *SystemServer) ReportComputerInfo(ip string, mac string) error {
	conn, err := grpc.Dial(server.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	name, _ := os.Hostname()
	defer conn.Close()
	client := pb.NewSystemManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.ReportServerinfo(ctx, &pb.ReportRequest{
		Ip:       ip,
		Mac:      mac,
		HostName: name,
	})
	if err != nil {
		return err
	}
	if response.Code != 0 {
		return fmt.Errorf("心跳失败:%s", response.Msg)
	}
	return nil
}
func InitSystemServer(address string) *SystemServer {
	return &SystemServer{
		address: address,
	}
}
