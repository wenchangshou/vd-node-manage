package engine

import (
	"context"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/layout"
	"time"

	"github.com/wenchangshou2/vd-node-manage/common/Event"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent-simple/service"
)

// Schedule 全局执行
type Schedule struct {
	ID             uint
	computerIp     string
	computerMac    string
	ComputerID     string
	register       bool
	lastReportTime time.Time
	serverInfo     *e.ServerInfo
	Count          int
	eventManage    *EventManage
	rpcAddress     string
	httpAddress    string
	serverFactory  *IService.ServiceFactory
	redisClient    *Event.RedisClient
	wm             layout.IManage
}

// 查询是否有新的分发任务
func (schedule Schedule) queryTask() {
	var (
		events []model.Event
		err    error
	)
	event := schedule.serverFactory.Event
	if events, err = event.QueryDeviceEvent(model.Initializes); err != nil {
		logging.GLogger.Warn(fmt.Sprintf("查询设备事件失败:%s", err.Error()))
		return
	}
	if len(events) > 0 {
		logging.GLogger.Info(fmt.Sprintf("当前有新的事件需要处理:%v", events))
		schedule.eventManage.AddWaitExecuteEvent(events)
	}
}

// loop 循环调度
func (schedule *Schedule) loop() {
	heartbeatTick := time.NewTicker(3 * time.Second)
	resourceDistributionTick := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-heartbeatTick.C:
			//schedule.ComputerService.Heartbeat()
		//case <-taskTick.C:
		//	schedule.TaskLoop()
		case <-resourceDistributionTick.C:
			fmt.Println("query task")
			schedule.queryTask()
		}
	}

}
func (schedule *Schedule) DeviceEvent(channel string, message []byte) error {
	fmt.Printf("channel:%s,message:%s\n", channel, string(message))
	return nil
}
func (schedule *Schedule) Start() {
	go schedule.redisClient.Subscribe(context.TODO(), schedule.DeviceEvent, fmt.Sprintf("device-%d", schedule.ID))
	go schedule.loop()
}

// InitSchedule 初始化调度程序
func InitSchedule(conf *g.GlobalConfig) error {
	var (
		err           error
		serverFactory *IService.ServiceFactory
	)
	rpcClient := &g.SingleConnRpcClient{
		RpcServer: fmt.Sprintf(conf.Server.RpcAddress),
		Timeout:   time.Second,
	}
	if serverFactory, err = IService.NewServiceFactory("rpc", conf.Server.ID, rpcClient); err != nil {
		return err
	}
	redisClient := Event.NewRedisClient(conf.Server.RedisAddress, 0, "")
	eventManage, _ := NewTaskManage(int32(conf.Task.Count), conf.Server.HttpAddress, serverFactory)
	eventManage.Start()
	schedule := &Schedule{
		ID:            conf.Server.ID,
		httpAddress:   conf.Server.HttpAddress,
		rpcAddress:    conf.Server.RpcAddress,
		serverFactory: serverFactory,
		redisClient:   redisClient,
		wm:            layout.InitLayoutManage(),
		eventManage:   eventManage,
	}

	schedule.Start()

	return nil
}
