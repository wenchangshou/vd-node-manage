package engine

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/dto"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/executor"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent-simple/service"
	"sync"
	"sync/atomic"
	"time"
)

var ActionGroup map[int]*executor.IExecute
var (
	GTaskExecute *TaskManage
)

type TaskChangeEventInfo struct {
	ID     string `json:"id"`
	Status uint
}

// NewTaskManage @title NewTaskManage
// @description 初始化任务执行
// @param maxExecutorCount 允许同时执行任务数
// @param httpRequestUri http请求地址
// @param rpcRequestUri rpc请求地址
func NewTaskManage(count int32, httpRequest string, taskService IService.TaskService) (*TaskManage, error) {
	g := executor.GenerateExecutorFactoryFunc(taskService, httpRequest)
	GTaskExecute = &TaskManage{
		maxExecutorCount:    count,
		taskAddNotify:       make(chan dto.Task),
		taskService:         taskService,
		notifyExecuteChange: make(chan TaskChangeEventInfo),
		TaskStatusList:      NewTaskList(),
		generator:           &g,
	}
	return GTaskExecute, nil
}

func init() {
}

// TaskManage  任务管理
type TaskManage struct {
	maxExecutorCount int32 //最大执行数
	executorCount    int32 //当前执行部数
	waitCount        int32 //等待数

	processTask         sync.Map //正在处理的任务列表
	waitTask            sync.Map //等待任务队列
	errorTask           sync.Map //错误队列
	taskAddNotify       chan dto.Task
	notifyExecuteChange chan TaskChangeEventInfo
	cancelFuncMap       sync.Map //用来处理取消任务使用
	TaskStatusList      *TaskList
	ActiveProcessCount  uint
	httpRequestUri      string
	rpcRequestUri       string
	taskService         IService.TaskService
	computerService     IService.ComputerService
	sync.RWMutex
	generator *executor.GeneratorFunction
}
type TaskStatus int

const (
	WAIT TaskStatus = iota
	EXECUTE
	ERROR
	UNKNOWN
)

type TaskList struct {
	sync.RWMutex
	items map[uint]TaskStatus
}

func NewTaskList() *TaskList {
	return &TaskList{items: make(map[uint]TaskStatus)}

}

//Delete 删除元素
func (taskList *TaskList) Delete(id uint) {
	taskList.Lock()
	defer taskList.Unlock()
	delete(taskList.items, id)
}

// LoadAndDeleteByStatus  加载并且删除元素
func (taskList *TaskList) LoadAndDeleteByStatus(status TaskStatus) (uint, bool) {
	taskList.Lock()
	defer taskList.Unlock()
	for k, v := range taskList.items {
		if v == status {
			delete(taskList.items, k)
			return k, true
		}
	}
	return 0, false
}

// Store 存储元素
func (taskList *TaskList) Store(id uint, status TaskStatus) {
	taskList.Lock()
	defer taskList.Unlock()
	taskList.items[id] = status
}
func (taskList *TaskList) Get(id uint) TaskStatus {
	taskList.RLock()
	defer taskList.Unlock()
	item, ok := taskList.items[id]
	if !ok {
		return UNKNOWN
	}
	return item
}

// execute 执行器
func (manage *TaskManage) execute(task dto.Task) {

	// 将任务设置成执行状态
	err := manage.taskService.SetTaskStatus([]uint{task.ID}, executor.EXECUTE)
	if err != nil {
		logging.GLogger.Info("更新任务状态失败")
		//manage.taskService.SetTaskStatus([]uint{task.ID}, executor.ERROR)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	manage.cancelFuncMap.LoadOrStore(task.ID, cancel)
	group := NewTaskGroup(task, ctx, manage.generator)
	status := group.Start()
	select {
	case status := <-status:
		manage.taskService.SetTaskStatus([]uint{task.ID}, status)
		atomic.AddInt32(&manage.executorCount, -1)
		manage.TaskStatusList.Delete(task.ID)
	}

}

// Loop 循环调度
func (manage *TaskManage) Loop() {
	loopTicker := time.NewTicker(time.Second)
	for {
		select {
		case <-loopTicker.C:
			manage.wake()
		}
	}
}

// wake 唤醒一个任务
func (manage *TaskManage) wake() {
	var (
		ok   bool
		wc   int32
		task dto.Task
	)
	wc = atomic.LoadInt32(&manage.waitCount)
	// 如果没有等待的元素，直接退出
	if wc == 0 {
		return
	}
	// 如果没有空闲的执行者，诚直接不处理
	if manage.GetTaskExecuteLave() <= 0 {
		return
	}
	// 从任务队列中指定状态的元素
	id, ok := manage.TaskStatusList.LoadAndDeleteByStatus(WAIT)
	// 如果没有任务元素，直接返回
	if !ok {
		return
	}
	atomic.AddInt32(&manage.waitCount, -1)
	atomic.AddInt32(&manage.executorCount, 1)
	tmp, ok := manage.waitTask.LoadAndDelete(id)
	if !ok {
		return
	}
	task = tmp.(dto.Task)
	manage.processTask.Store(id, task)
	go manage.execute(task)
}
func (manage *TaskManage) AddTaskByExecuteList(task []dto.Task) {
	manage.Lock()
	defer manage.Unlock()
}

func (manage *TaskManage) AddTask(tasks []dto.Task) error {
	manage.Lock()
	defer manage.Unlock()
	for _, task := range tasks {
		atomic.AddInt32(&manage.waitCount, 1)
		manage.TaskStatusList.Store(task.ID, WAIT)
		manage.waitTask.LoadOrStore(task.ID, task)
	}
	return nil
}
func (manage *TaskManage) Start() {
	go manage.Loop()
}

// GetTaskExecuteLave 获取任务空闲数
func (manage *TaskManage) GetTaskExecuteLave() int32 {
	count := atomic.LoadInt32(&manage.executorCount)
	return manage.maxExecutorCount - count
}
