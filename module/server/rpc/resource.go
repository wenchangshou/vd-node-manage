package rpc

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	model2 "github.com/wenchangshou2/vd-node-manage/module/server/model"
)

func (resource Resource) Query(args *model.ResourceQueryRequest, _ *model.ResourceQueryResponse) error {
	model2.GetResourceById(args.ID)
	return nil

}
