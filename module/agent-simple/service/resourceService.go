package IService

import "github.com/wenchangshou2/vd-node-manage/common/model"

type ResourceService interface {
	QueryResource(id uint) (model.ResourceInfo, error)
}
