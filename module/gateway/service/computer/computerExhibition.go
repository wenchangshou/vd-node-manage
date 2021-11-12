package computer

import (
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
)

type ComputerExhibitionOpenService struct {
	ComputerID   int `json:"id" uri:"id"`
	ExhibitionID int `json:"exhibitionID"`
}

func (service *ComputerExhibitionOpenService) Open() serializer.Response {
	return serializer.Response{}
}
