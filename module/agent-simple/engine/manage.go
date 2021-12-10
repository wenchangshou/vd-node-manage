package engine

import (
	"context"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/common/model"
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
func NewTaskManage(count int32, httpRequest string, eventService IService.EventService) (*TaskManage, error) {
	g := executor.GenerateExecutorFactoryFunc(eventService, httpRequest)
	GTaskExecute = &TaskManage{
		maxExecutorCount:    count,
		taskAddNotify:       make(chan dto.Task),
		eventService:        eventService,
		notifyExecuteChange: make(chan TaskChangeEventInfo),
		EventStatusList:     NewTaskList(),
		generator:           g,
	}
	return GTaskExecute, nil
}

func init() {
}

// TaskManage  任务管理
type TaskManage struct {
	maxExecutorCount    int32    //最大执行数
	executorCount       int32    //当前执行部数
	waitCount           int32    //等待数
	processTask         sync.Map //正在处理的任务列表
	waitTask            sync.Map //等待任务队列
	errorTask           sync.Map //错误队列
	taskAddNotify       chan dto.Task
	notifyExecuteChange chan TaskChangeEventInfo
	cancelFuncMap       sync.Map //用来处理取消任务使用
	EventStatusList     *EventList
	ActiveProcessCount  uint
	httpRequestUri      string
	rpcRequestUri       string
	eventService        IService.EventService
	computerService     IService.ComputerService
	sync.RWMutex
	generator executor.GeneratorFunction
}

type EventList struct {
	sync.RWMutex
	items map[uint]model.EventStatus
}

func NewTaskList() *EventList {
	return &EventList{items: make(map[uint]model.EventStatus)}

}

//Delete 删除元素
func (eventList *EventList) Delete(id uint) {
	eventList.Lock()
	defer eventList.Unlock()
	delete(eventList.items, id)
}

// LoadAndDeleteByStatus  加载并且删除元素
func (eventList *EventList) LoadAndDeleteByStatus(status model.EventStatus) (uint, bool) {
	eventList.Lock()
	defer eventList.Unlock()
	for k, v := range eventList.items {
		if v == status {
			delete(eventList.items, k)
			return k, true
		}
	}
	return 0, false
}

// Store 存储元素
func (eventList *EventList) Store(id uint, status model.EventStatus) {
	eventList.Lock()
	defer eventList.Unlock()
	eventList.items[id] = status
}
func (eventList *EventList) Get(id uint) model.EventStatus {
	eventList.RLock()
	defer eventList.Unlock()
	item, ok := eventList.items[id]
	if !ok {
		return model.UNKNOWN
	}
	return item
}

// execute 执行器
func (manage *TaskManage) execute(event model.Event) {
	// 将任务设置成执行状态
	err := manage.eventService.SetEventStatus([]uint{event.ID}, executor.EXECUTE)
	if err != nil {
		logging.GLogger.Info("更新任务状态失败")
		manage.eventService.SetEventStatus([]uint{event.ID}, executor.ERROR)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	manage.cancelFuncMap.LoadOrStore(event.ID, cancel)
	fmt.Println(ctx)
	m := NewEventExecuteManage(event, ctx, manage.generator)
	status := m.Start()
	select {
	case s := <-status:
		manage.eventService.SetEventStatus([]uint{event.ID}, s)
		atomic.AddInt32(&manage.executorCount, -1)
		manage.EventStatusList.Delete(event.ID)
	}
	//e,err:=manage.generator(event.Action, event.ID, event.Params)
	//group := NewEventExecuteManage(task, ctx, manage.generator)
	//status := group.Start()
	//select {
	//case status := <-status:
	//	manage.taskService.SetTaskStatus([]uint{task.ID}, status)
	//	atomic.AddInt32(&manage.executorCount, -1)
	//	manage.EventStatusList.Delete(task.ID)
	//}

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
		ok    bool
		wc    int32
		event model.Event
		id    uint
		tmp   interface{}
	)
	// 如果没有等待的任务或者可执行的任务为0
	if wc = atomic.LoadInt32(&manage.waitCount); wc == 0 || manage.GetTaskExecuteLave() <= 0 {
		return
	}
	// 从任务队列中指定状态的元素
	if id, ok = manage.EventStatusList.LoadAndDeleteByStatus(model.WAITING); !ok {
		return
	}
	atomic.AddInt32(&manage.waitCount, -1)
	atomic.AddInt32(&manage.executorCount, 1)
	if tmp, ok = manage.waitTask.LoadAndDelete(id); !ok {
		return
	}
	event = tmp.(model.Event)
	manage.processTask.Store(id, event)
	go manage.execute(event)
}
func (manage *TaskManage) AddTaskByExecuteList(task []dto.Task) {
	manage.Lock()
	defer manage.Unlock()
}

// AddWaitExecuteEvent 添加任务到等待列表中
func (manage *TaskManage) AddWaitExecuteEvent(events []model.Event) error {
	manage.Lock()
	defer manage.Unlock()
	for _, event := range events {
		atomic.AddInt32(&manage.waitCount, 1)
		manage.EventStatusList.Store(event.ID, model.WAITING)
		manage.waitTask.LoadOrStore(event.ID, event)
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
