package rpc

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
)

type DeviceRpcService struct {
	ID     uint `json:"id"`
	Client *g.SingleConnRpcClient
}

func (service DeviceRpcService) Heartbeat() error {
	//TODO implement me
	panic("implement me")
}

func (service DeviceRpcService) AddComputerProject(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (service DeviceRpcService) DeleteComputerProject(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (service DeviceRpcService) ReportServiceInfo(id uint, ip string, mac string, name string) error {
	return nil
}
func (service DeviceRpcService) Report() error {
	return nil
}
func (service DeviceRpcService) IsRegister() (bool, error) {
	return true, nil
}
func (service DeviceRpcService) AddComputerResource(resourceID uint) error {
	var (
		err error
	)
	req := model.DeviceAddResourceRequest{
		ID:         service.ID,
		ResourceID: resourceID,
	}
	reply := model.SimpleRpcResponse{}
	if err = service.Client.Call("Device.AddDeviceResource", &req, &reply); err != nil {
		return err
	}
	return nil
}
func (service DeviceRpcService) DeleteComputerResource(resourceID uint) error {
	return nil
}

// NewDeviceRpcService 新的设备rpc服务
func NewDeviceRpcService(id uint, client *g.SingleConnRpcClient) *DeviceRpcService {
	return &DeviceRpcService{
		ID:     id,
		Client: client,
	}
}
