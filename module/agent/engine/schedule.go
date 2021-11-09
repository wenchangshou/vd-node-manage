package engine

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent/engine/executor"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent/service"
	HttpService "github.com/wenchangshou2/vd-node-manage/module/agent/service/impl"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/logging"
	"time"
	"github.com/pkg/errors"
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
		logging.G_Logger.Warn(fmt.Sprintf("调用获取任务接口失败:%s", err.Error()))
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
	taskTick := time.NewTicker(5*time.Second)
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
func InitSchedule(s *e.ServerInfo, id string) error {
	computerHttpService := HttpService.NewComputerHttpService(id, s.Ip, s.Port)
	taskService := HttpService.NewTaskHttpService(id, s.Ip, s.Port)
	taskManage, err := NewTaskManage(int32(conf.TaskConfig.Count), fmt.Sprintf("%s:%d",s.Ip,s.Port),taskService, computerHttpService)
	if err != nil {
		return errors.Wrap(err, "创建任务管理器失败")
	}
	taskManage.Start()
	schedule := &Schedule{
		ComputerService: computerHttpService,
		TaskService:     taskService,
		taskManage:      taskManage,
		serverInfo:      s,
	}

	schedule.Start()

	return nil
}
