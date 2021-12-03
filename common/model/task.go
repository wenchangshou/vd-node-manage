package model

import "github.com/wenchangshou2/vd-node-manage/module/agent-simple/dto"

const (
	TaskInitialization = iota
	TaskProcess
	TASKError
)

type QueryDeviceResourceDistributionRequest struct {
	DeviceID uint `json:"device_id"`
	Tasks    []dto.Task `json:"tasks"`
}
type QueryDeviceResourceDistributionResponse struct {
	DeviceID uint `json:"device_id"`
	Count    int
	Tasks     []dto.Task
}
