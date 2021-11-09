package rpcServer

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ResourceServer struct {
	pb.UnimplementedResourceManagementServer
}

func (server *ResourceServer) GetResourceDetailedInfo(ctx context.Context, id *wrapperspb.StringValue) (response *pb.GetResourceDetailedInfoResponse, err error) {
	resource, err := model.GetResourceById(id.GetValue())
	if err != nil {
		return nil, err
	}
	file := resource.File

	response = &pb.GetResourceDetailedInfoResponse{
		ID:       resource.ID,
		Name:     resource.Name,
		Category: resource.Category,
		File: &pb.GetResourceDetailedInfoResponse_File{
			ID:         file.ID,
			Name:       file.Name,
			Size:       int64(file.Size),
			Url:        "upload/" + file.SourceName,
			SourceName: file.SourceName,
			Uuid:       file.Uuid,
		},
	}
	return response, nil
}
