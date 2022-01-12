package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/wenchangshou/vd-node-manage/common/cache"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g/db"
	bolt "go.etcd.io/bbolt"

	"github.com/wenchangshou/vd-node-manage/module/agent-simple/engine/layout"

	"github.com/wenchangshou/vd-node-manage/common/Event"
	"github.com/wenchangshou/vd-node-manage/common/logging"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
	IService "github.com/wenchangshou/vd-node-manage/module/agent-simple/service"
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
func (schedule *Schedule) openLayout(req model.EventRequest) model.EventReply {
	var (
		err error
	)
	reply := model.EventReply{EventID: req.EventID}
	args := model.OpenLayoutCmdParams{}
	if err := json.Unmarshal(req.Arguments, &args); err != nil {
		reply.Err = err
		reply.Msg = "解析json错误"
		return reply
	}
	if err = schedule.layoutManage.OpenLayout(args); err != nil {
		reply.Err = err
		reply.Msg = "打开布局失败"
		return reply
	}
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}
func (schedule *Schedule) closeLayout(req model.EventRequest) model.EventReply {
	var (
		err error
	)
	reply := model.EventReply{EventID: req.EventID}
	if err = schedule.layoutManage.CloseLayout(); err != nil {
		reply.Err = err
		reply.Msg = "关闭布局失败"
		return reply
	}
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}

func (schedule Schedule) CheckResourceExists(req model.EventRequest) model.EventReply {
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}
func (schedule *Schedule) control(req model.EventRequest) model.EventReply {
	reply := model.EventReply{
		EventID: req.EventID,
	}
	m := schedule.layoutManage
	args := model.ControlWindowCmdParams{}
	err := json.Unmarshal(req.Arguments, &args)
	if err != nil {
		reply.Err = err
		reply.Msg = "解析控制命令失败"
		return reply
	}
	err = m.Control(args, false)
	if err != nil {
		reply.Err = err
		reply.Msg = "控制窗口失败"
		return reply
	}
	return model.GenerateSimpleSuccessEventReply(req.EventID)
}

// DeviceEvent 接收服务端事件
func (schedule *Schedule) DeviceEvent(_ string, message []byte) (r []byte, err error) {
	req := model.EventRequest{}
	reply := model.EventReply{}
	if err = json.Unmarshal(message, &req); err != nil {
		return
	}
	if req.Action == "openLayout" {
		reply = schedule.openLayout(req)
	} else if req.Action == "closeLayout" {
		reply = schedule.closeLayout(req)
	} else if req.Action == "control" {
		reply = schedule.control(req)
	} else if req.Action == "checkResourceExists" {
		reply = schedule.CheckResourceExists(req)
	}
	if !req.Reply {
		return nil, nil
	}
	b, _ := json.Marshal(reply)
	return b, nil
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
	return schedule.layoutManage.OpenLayout(cmd)
}
func (schedule *Schedule) Start() {
	schedule.Startup()
	go schedule.redisClient.Subscribe(context.TODO(), schedule.DeviceEvent, fmt.Sprintf("device-%d", schedule.ID))
	go schedule.loop()
}

// InitSchedule 初始化调度程序
func InitSchedule(conf *g.GlobalConfig, driver *cache.Driver) error {
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
		return err
	}
	redisClient := Event.NewRedisClient(serverInfo.Redis.Address, 0, "")
	eventManage, _ := NewEventManage(int32(conf.Task.Count), driver, db.GDB, serverInfo.Http.Address, serverFactory)
	eventManage.Start()
	schedule := &Schedule{
		ID:            serverInfo.ID,
		httpAddress:   serverInfo.Http.Address,
		rpcAddress:    serverInfo.Rpc.Address,
		serverFactory: serverFactory,
		redisClient:   redisClient,
		layoutManage:  layout.InitLayoutManage(db.GDB),
		eventManage:   eventManage,
		cacheDriver:   driver,
		db:            db.GDB,
	}
	schedule.Start()
	return nil
}
