package engine

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/executor"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent-simple/service"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/service/impl/http"
	"time"
)

// Schedule 全局执行
type Schedule struct {
	computerIp      string
	computerMac     string
	ComputerID      string
	register        bool
	lastReportTime  time.Time
	ComputerService IService.ComputerService
	TaskService     IService.TaskService
	serverInfo      *e.ServerInfo
	Count           int
	taskManage      *TaskManage
	RpcAddress      string
	HttpAddress     string
}

func (schedule Schedule) TaskLoop() {
	// 判断当前任务管理器是否有空闲的资源
	idleCount := schedule.taskManage.GetTaskExecuteLave()
	if idleCount <= 0 {
		return
	}
	//  获取根据剩余的资源来查询任务数
	tasks, err := schedule.TaskService.GetTasks(executor.INITIALIZE, int(idleCount))
	if err != nil {
		logging.GLogger.Warn(fmt.Sprintf("调用获取任务接口失败:%s", err.Error()))
		return
	}
	if len(tasks) == 0 {
		return
	}

	// 将任务管理器添加新的任务
	schedule.taskManage.AddTask(tasks)
}
func (schedule *Schedule) loop() {
	heartbeatTick := time.NewTicker(3 * time.Second)
	taskTick := time.NewTicker(5 * time.Second)
	schedule.ComputerService.Report()
	for {
		select {
		case <-heartbeatTick.C:
			schedule.ComputerService.Heartbeat()
		case <-taskTick.C:
			schedule.TaskLoop()
		}
	}

}
func (schedule *Schedule) Start() {
	go schedule.loop()
}

// InitSchedule 初始化调度程序
func InitSchedule(httpAddress string,rpcAddress string, id string) error {
	cfg:=g.Config()
	computerHttpService := http.NewComputerHttpService(id,httpAddress)
	taskService := http.NewTaskHttpService(id, httpAddress)
	taskManage, err := NewTaskManage(int32(cfg.Task.Count),  cfg.Server.HttpAddress, taskService, computerHttpService)
	if err != nil {
		return errors.Wrap(err, "创建任务管理器失败")
	}
	taskManage.Start()
	schedule := &Schedule{
		ComputerService: computerHttpService,
		TaskService:     taskService,
		taskManage:      taskManage,
		HttpAddress:httpAddress,
		RpcAddress:rpcAddress,
	}

	schedule.Start()

	return nil
}