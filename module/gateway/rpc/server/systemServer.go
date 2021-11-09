package rpcServer

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server/pb"
)

type SystemServer struct {
	pb.UnimplementedSystemManagementServer
}

func (server *SystemServer) ReportServerinfo(ctx context.Context, request *pb.ReportRequest) (response *pb.ReportResponse, err error) {
	client := model.Computer{
		Ip:       request.Ip,
		Mac:      request.Mac,
		HostName: request.HostName,
	}
	if client.IsExistByMac() {
		err = client.UpdateByMac()
		return &pb.ReportResponse{
			Code: 0,
			Msg:  "成功",
		}, err
	}
	err = client.Create()
	return &pb.ReportResponse{}, err
}
