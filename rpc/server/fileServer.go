package rpcServer

import (
	"context"
	"fmt"

	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/rpc/server/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type FileServer struct {
	pb.UnimplementedFileManagementServer
}

func (server *FileServer) GetFileInfoByProjectReleaseID(ctx context.Context, id *wrapperspb.StringValue) (*pb.GetFileInfoByProjectReleaseIDResponse, error) {
	pr, err := model.GetProjectReleaseByID(id.GetValue())
	if err != nil {
		return nil, err
	}
	result := pb.GetFileInfoByProjectReleaseIDResponse{
		ID:         pr.File.ID,
		Size:       int64(pr.File.Size),
		Url:        "upload/" + pr.File.SourceName,
		SourceName: pr.File.SourceName,
		Name:       pr.File.Name,
		Uuid:       pr.File.Uuid,
	}
	fmt.Printf("file info:%v\n", pr.File)
	return &result, nil
}
