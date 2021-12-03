package layout

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/zebus"
)

type Layout struct {
}
type Style struct {
	WindowStyle string `json:"window_style"`
	Theme       string `json:"theme"`
}
type Source struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Fname string `json:"fname"`
	URI   string `json:"uri"`
}
type Window struct {
	ID        string `json:"id"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Z         int    `json:"z"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Service   string `json:"service"`
	Arguments string `json:"arguments"`
	Style     Style  `json:"style"`
}
type OpenMultiScreenService struct {
	LayoutID string   `json:"layout_id"`
	Kill     bool     `json:"kill"`
	Style    Style    `json:"style"`
	Windows  []Window `json:"windows"`
}

// type LayoutOpenMultiScreenService struct {
// 	Action    string               `json:"action"`
// 	Arguments OpenMultiScreenParam `json:"arguments"`
// }
type OpenMultiScreenResultFrom struct {
}

func (service *OpenMultiScreenService) Open(computerId int) serializer.Response {
	var result *OpenMultiScreenResultFrom
	computer, err := model.GetComputerById(computerId)
	if err != nil {
		return serializer.Err(serializer.CodeNotFindComputerErr, "没有找到对应的计算机", err)
	}
	serverInfo, err := zebus.G_Zebus.GetClients()
	if err != nil {
		return serializer.Err(serializer.CodeCallZebusApiErr, "调用zebus接口错误", err)
	}
	if !serverInfo.IsExistServer(computer.Ip, "vd-ResourcesService") {
		return serializer.Err(serializer.CodeNotFindDeviceErr, "目标计算机服务未在线", nil)
	}
	topic := fmt.Sprintf("/zebus/%s/vd-ResourcesService", computer.Ip)
	b, err := json.Marshal(service)
	if err != nil {
		return serializer.Err(serializer.CodeJsonUnMarkshalErr, "解析body失败", err)
	}

	err = zebus.G_Zebus.PutV2(topic, string(b), 0, 0, result)
	if err != nil {
		return serializer.Err(serializer.CodeSendZebusMessageErr, "推送zebus消息失败", err)
	}
	return serializer.Response{}
}
