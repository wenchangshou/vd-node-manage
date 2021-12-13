package engine

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/wenchangshou2/vd-node-manage/common/Event"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent-simple/service"
	"time"
)

// Schedule 全局执行
type Schedule struct {
	ID             uint
	computerIp     string
	computerMac    string
	ComputerID     string
	register       bool
	lastReportTime time.Time
	TaskService    IService.TaskService
	serverInfo     *e.ServerInfo
	Count          int
	taskManage     *TaskManage
	RpcAddress     string
	HttpAddress    string
	ServerFactory  *IService.ServiceFactory
	redisClient    *Event.RedisClient
}

// 查询是否有新的分发任务
func (schedule Schedule) queryTask() {
	event := schedule.ServerFactory.Event
	tasks, err := event.QueryDeviceEvent(model.Initializes)
	if len(tasks) > 0 {
		logging.GLogger.Info(fmt.Sprintf("当前有新的事件需要处理:%v", tasks))
		schedule.taskManage.AddWaitExecuteEvent(tasks)
	}
	fmt.Println("tasks", tasks, err)
}

// loop 循环调度
func (schedule *Schedule) loop() {
	heartbeatTick := time.NewTicker(3 * time.Second)
	//taskTick := time.NewTicker(5 * time.Second)
	resourceDistributionTick := time.NewTicker(5 * time.Second)
	//schedule.ComputerService.Report()
	for {
		select {
		case <-heartbeatTick.C:
			//schedule.ComputerService.Heartbeat()
		//case <-taskTick.C:
		//	schedule.TaskLoop()
		case <-resourceDistributionTick.C:
			schedule.queryTask()
		}
	}

}
func (schedule *Schedule) DeviceEvent(channel string, message []byte) error {

}
func (schedule *Schedule) Start() {
	schedule.redisClient.Subscribe(context.TODO(), schedule.DeviceEvent, fmt.Sprintf("device-%d", schedule.ID))
	go schedule.loop()
}

// InitSchedule 初始化调度程序
func InitSchedule(httpAddress string, rpcAddress string, id uint) error {
	cfg := g.Config()
	rpcClient := &g.SingleConnRpcClient{
		RpcServer: fmt.Sprintf(cfg.Server.RpcAddress),
		Timeout:   time.Second,
	}
	serverFactory, err := IService.NewServiceFactory("rpc", id, rpcClient)
	if err != nil {
		return err
	}
	//taskService := rpc.NewTaskRpcService(id)
	//eventService := rpc.NewEventRpcService(id, rpcClient)
	taskManage, err := NewTaskManage(int32(cfg.Task.Count), cfg.Server.HttpAddress, serverFactory)
	if err != nil {
		return errors.Wrap(err, "创建任务管理器失败")
	}
	taskManage.Start()
	redisClient := Event.NewRedisClient("192.168.10.31:30000", 0, "")
	schedule := &Schedule{
		ID:            id,
		taskManage:    taskManage,
		HttpAddress:   httpAddress,
		RpcAddress:    rpcAddress,
		ServerFactory: serverFactory,
		redisClient:   redisClient,
	}

	schedule.Start()

	return nil
}
