package IService

import "github.com/wenchangshou/vd-node-manage/common/model"

type ProjectService interface {
	QueryProject(id uint) (*model.ProjectInfo, error)
}
