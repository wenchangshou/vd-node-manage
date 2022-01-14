package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/module/core/g/db"
	bolt "go.etcd.io/bbolt"

	"github.com/wenchangshou/vd-node-manage/module/core/engine/layout"

	"github.com/wenchangshou/vd-node-manage/common/Event"
	"github.com/wenchangshou/vd-node-manage/common/logging"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
	IService "github.com/wenchangshou/vd-node-manage/module/core/service"
)

// Schedule 全局执行
type Schedule struct {
	ID             uint
	computerIp     string
	computerMac    string
	ComputerID     string
	register       bool
	lastReportTime time.Time
	Count          int
	eventManage    *EventManage
	rpcAddress     string
	httpAddress    string
	serverFactory  *IService.ServiceFactory
	redisClient    *Event.RedisClient
	layoutManage   layout.IManage
	cacheDriver    *cache.Driver
	db             *bolt.DB
	threadMap      map[string]uint32
}

// queryTask 查询是否有新的分发任务
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
	playerActiveInfoTick := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-heartbeatTick.C:
		case <-playerActiveInfoTick.C:
			schedule.reportPlayerRunInfo()
		case <-resourceDistributionTick.C:
			schedule.queryTask()
		}
	}

}
func (schedule *Schedule) reportPlayerRunInfo() {
	if schedule.layoutManage.GetLayoutID() == "" {
		return
	}
	res, err := schedule.layoutManage.GetLayoutRunInfo()
	if err != nil {
		return
	}
	b, _ := json.Marshal(res)
	cache.Set(fmt.Sprintf("device-%d-%s", schedule.ID, schedule.layoutManage.GetLayoutID()), string(b), 10)
}

func (schedule *Schedule) Startup() error {
	cmd := model.OpenLayoutCmdParams{}
	startup, err := schedule.serverFactory.Device.GetDeviceStartup(schedule.ID)
	if err != nil {
		logging.GLogger.Warn("请求开机启动项失败:" + err.Error())
		return err
	}
	if startup == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(startup), &cmd); err != nil {
		logging.GLogger.Warn("解析开机命令失败:" + err.Error())
	}
	schedule.threadMap, err = schedule.layoutManage.OpenLayout(cmd)
	return err
}
func (schedule *Schedule) Start() {
	schedule.Startup()
	go schedule.loop()
}
func (schedule *Schedule) InitEventManage() error {
	serverInfo := g.GetServerInfo()
	em, err := Event.NewEvent(serverInfo.Event.Provider, serverInfo.Event.Arguments)
	if err != nil {
		return err
	}
	go em.Subscribe(context.TODO(), schedule.DeviceEvent, fmt.Sprintf("device-%d", schedule.ID))
	return nil
}

func (schedule Schedule) Exit() {
	schedule.layoutManage.CloseLayout()
}

// InitSchedule 初始化调度程序
func InitSchedule(conf *g.GlobalConfig, driver *cache.Driver) (*Schedule, error) {
	var (
		err           error
		serverFactory *IService.ServiceFactory
	)
	serverInfo := g.GetServerInfo()
	rpcClient := &g.SingleConnRpcClient{
		RpcServer: fmt.Sprintf(serverInfo.Rpc.Address),
		Timeout:   time.Second,
	}

	if serverFactory, err = IService.NewServiceFactory("rpc", serverInfo.ID, rpcClient); err != nil {
		return nil, err
	}
	eventManage, _ := NewEventManage(int32(conf.Task.Count), driver, db.GDB, serverInfo.Http.Address, serverFactory)
	eventManage.Start()
	schedule := &Schedule{
		ID:            serverInfo.ID,
		httpAddress:   serverInfo.Http.Address,
		rpcAddress:    serverInfo.Rpc.Address,
		serverFactory: serverFactory,
		layoutManage:  layout.InitLayoutManage(db.GDB),
		eventManage:   eventManage,
		cacheDriver:   driver,
		db:            db.GDB,
	}
	schedule.InitEventManage()
	schedule.Start()
	return schedule, nil
}
