package rpc

import (
	"context"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ResourceServer struct {
	pb.UnimplementedResourceManagementServer
}

func (server *ResourceServer) GetResourceDetailedInfo(ctx context.Context, id *wrapperspb.UInt32Value) (response *pb.GetResourceDetailedInfoResponse, err error) {
	resource, err := model.GetResourceById(uint(id.GetValue()))
	if err != nil {
		return nil, err
	}
	file := resource.File

	response = &pb.GetResourceDetailedInfoResponse{
		ID:       uint32(resource.ID),
		Name:     resource.Name,
		Category: resource.Category,
		File: &pb.GetResourceDetailedInfoResponse_File{
			ID:         int32(file.ID),
			Name:       file.Name,
			Size:       int64(file.Size),
			Url:        "upload/" + file.SourceName,
			SourceName: file.SourceName,
			Uuid:       file.Uuid,
		},
	}
	return response, nil
}
