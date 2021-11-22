package cron

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"time"
)

// ReportDeviceStatus 上传设备状态
func ReportDeviceStatus() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go reportAgentStatus(time.Duration(g.Config().Heartbeat.Interval) * time.Second)
	}
}
func reportAgentStatus(interval time.Duration) {
	for {
		hostname, err := g.Hostname()
		if err != nil {
			hostname = fmt.Sprintf("error:%s", err.Error())
		}
		req := model.DeviceReportRequest{
			ID:       "",
			Hostname: hostname,
			Ip:       g.IP(),
		}
		fmt.Println("req", req)
		time.Sleep(interval)
	}
}
