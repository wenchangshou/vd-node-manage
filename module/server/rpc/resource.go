package rpc

import (
	"github.com/wenchangshou/vd-node-manage/common/model"
	model2 "github.com/wenchangshou/vd-node-manage/module/server/model"
)

func (resource Resource) Query(request *model.ResourceQueryRequest, response *model.ResourceQueryResponse) error {
	var (
		r   *model2.Resource
		err error
	)
	if r, err = model2.GetResourceById(request.ID); err != nil {
		response.Code = 400
		response.Msg = err.Error()
		return err
	}
	response.Resource = model.ResourceInfo{
		ID:      r.ID,
		Name:    r.Name,
		Uri:     r.Uri,
		Service: r.Service,
		Status:  r.Status,
		Md5:     r.Md5,
	}
	return nil
}
