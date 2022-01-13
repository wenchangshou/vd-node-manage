package cron

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
	"log"
	"sync/atomic"
	"time"
)

var reportFlag int64 = 0

// ReportDeviceStatus 上传设备状态
func ReportDeviceStatus() {

	if reportFlag != 1 && g.GetServerInfo().Rpc.Address != "" {
		atomic.StoreInt64(&reportFlag, 1)
		go reportAgentStatus(time.Second)
	}
}

//reportAgentStatus  上报节点状态
func reportAgentStatus(interval time.Duration) {
	for {
		r := make(map[string]interface{})
		hostname, err := g.Hostname()
		hid := g.GetServerInfo().ID
		if err != nil {
			hostname = fmt.Sprintf("error:%s", err.Error())
		}
		r["hostname"] = hostname
		r["ip"] = g.IP()
		b, _ := json.Marshal(r)
		req := model.DeviceReportRequest{
			ID:   hid,
			Info: string(b),
		}
		var resp model.SimpleRpcResponse
		err = g.ServerRpcClient.Call("Device.ReportStatus", req, &resp)
		if err != nil || resp.Code != 0 {
			log.Println("Call Device.ReportStatus fail", err, "Request:", req, "Response:", resp)
		}
		time.Sleep(interval)
	}
}