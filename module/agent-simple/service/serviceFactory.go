package IService

import (
	"errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/service/impl/rpc"
)

type ServiceFactory struct {
	Device   DeviceService
	Event    EventService
	Resource ResourceService
}

func NewServiceFactory(protocol string, id uint, client *g.SingleConnRpcClient) (*ServiceFactory, error) {
	s := &ServiceFactory{}
	if protocol == "rpc" {
		event := rpc.NewEventRpcService(id, client)
		device := rpc.NewDeviceRpcService(id, client)
		s.Event = event
		s.Device = device
		return s, nil
	}
	return nil, errors.New("未找到對應的協議")

}
