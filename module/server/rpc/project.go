package rpc

import (
	"github.com/wenchangshou/vd-node-manage/common/model"
	model2 "github.com/wenchangshou/vd-node-manage/module/server/model"
)

func (project Project) Query(request *model.NormalIdRequest, response *model.ProjectQueryResponse) error {
	var (
		r   *model2.Project
		err error
	)
	if r, err = model2.GetProjectByID(request.ID); err != nil {
		response.Code = 400
		response.Msg = err.Error()
		return err
	}
	response.Project = model.ProjectInfo{
		ID:      r.ID,
		Name:    r.Name,
		Uri:     r.Uri,
		Service: r.Service,
		Status:  r.Status,
		Startup: r.Startup,
		Md5:     r.Md5,
	}
	return nil
}
