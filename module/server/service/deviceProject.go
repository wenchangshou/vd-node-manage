package service

import (
	model2 "github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/common/serializer"
	"github.com/wenchangshou/vd-node-manage/module/server/model"
)

type ProjectForm struct {
	ID      uint   `json:"id"`
	Md5     string `json:"md5"`
	URI     string `json:"uri"`
	Name    string `json:"name"`
	Service string `json:"service"`
	Startup string `json:"startup"`
}

type DeviceProjectAddService struct {
	ID       uint          `json:"id" uri:"id"`
	Projects []ProjectForm `json:"projects"`
}

func (service DeviceProjectAddService) Add() serializer.Response {
	var (
		err error
	)
	IDMap := make([]IDRelation, 0)
	for _, v := range service.Projects {
		var _project *model.Project
		_project, err = service.add(v)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "添加项目失败", err)
		}
		m := model.Event{
			DeviceID: service.ID,
			Active:   false,
			Action:   model2.InstallProjectAction,
			Status:   model2.Initializes,
		}
		m.ProjectId = _project.ID
		if err = m.Add(); err != nil {
			return serializer.Err(serializer.CodeDBError, "添加项目分发事件失败", err)
		}
		r := IDRelation{
			SId: v.ID,
			DId: _project.ID,
		}
		IDMap = append(IDMap, r)
	}
	return serializer.Response{
		Data: map[string]interface{}{
			"relation": IDMap,
		},
	}
}

func (service DeviceProjectAddService) add(v ProjectForm) (project *model.Project, err error) {
	project = &model.Project{
		Name:    v.Name,
		Service: v.Service,
		Uri:     v.URI,
		Status:  0,
		Md5:     v.Md5,
		Startup: v.Startup,
	}
	_, err = project.Create()
	return
}
