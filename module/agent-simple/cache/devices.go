package cache

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"sync"
)

type SafeDevices struct {
	sync.RWMutex
	M map[string]*model.DeviceUpdateInfo
}

var Devices = NewSafeDevices()

func NewSafeDevices() *SafeDevices {
	return &SafeDevices{M: make(map[string]*model.DeviceUpdateInfo)}
}
func (this *SafeDevices) Put(req *model.DeviceReportRequest) {
}
func (this *SafeDevices) Get(hid string) (*model.DeviceUpdateInfo, bool) {
	this.RLock()
	defer this.RUnlock()
	val, exists := this.M[hid]
	return val, exists
}
