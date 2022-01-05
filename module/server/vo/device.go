package vo

import "github.com/wenchangshou/vd-node-manage/module/server/model"

type DeviceVo struct {
	ID             uint   `json:"id"`
	Code           string `json:"code"`
	ConnType       string `json:"connType"`
	Name           string `json:"name"`
	HostName       string `json:"hostName"`
	Status         int    `json:"status"`
	RegionID       int    `json:"regionId"`
	Online         bool   `json:"online"`
	Detailed       string `json:"detailed"`
	LastOnlineTime int64  `json:"lastOnlineTime"`
}

func DeviceDoToVo(d *model.Device) *DeviceVo {
	if d == nil {
		return nil
	}
	return &DeviceVo{
		ID:       d.ID,
		Code:     d.Code,
		ConnType: d.ConnType,
		Name:     d.Name,
		HostName: d.HostName,
		Status:   d.Status,
		RegionID: d.RegionId,
	}

}
