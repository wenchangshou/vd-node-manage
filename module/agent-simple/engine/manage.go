package engine

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/dto"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/executor"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent-simple/service"
)

var ActionGroup map[int]*executor.IExecute
var (
	GTaskExecute *EventManage
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
func NewTaskManage(count int32, httpRequest string, service *IService.ServiceFactory) (*EventManage, error) {
	g := executor.GenerateExecutorFactoryFunc(service, httpRequest)
	GTaskExecute = &EventManage{
		maxExecutorCount:    count,
		taskAddNotify:       make(chan dto.Task),
		ServerFactory:       service,
		notifyExecuteChange: make(chan TaskChangeEventInfo),
		EventStatusList:     NewTaskList(),
		generator:           g,
	}
	return GTaskExecute, nil
}

func init() {
}

// EventManage  任务管理
type EventManage struct {
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
	sync.RWMutex
	generator     executor.GeneratorFunction
	ServerFactory *IService.ServiceFactory
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
func (manage *EventManage) execute(event model.Event) {
	e := manage.ServerFactory.Event
	// 将任务设置成执行状态
	err := e.SetEventStatus([]uint{event.ID}, executor.EXECUTE)
	if err != nil {
		logging.GLogger.Info("更新任务状态失败")
		e.SetEventStatus([]uint{event.ID}, executor.ERROR)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	manage.cancelFuncMap.LoadOrStore(event.ID, cancel)
	m := NewEventExecuteManage(event, ctx, manage.generator)
	status := m.Start()
	select {
	case s := <-status:
		e.SetEventStatus([]uint{event.ID}, s)
		atomic.AddInt32(&manage.executorCount, -1)
		manage.EventStatusList.Delete(event.ID)
	}

}

// Loop 循环调度
func (manage *EventManage) Loop() {
	loopTicker := time.NewTicker(time.Second)
	for {
		select {
		case <-loopTicker.C:
			manage.wake()
		}
	}
}

// wake 唤醒一个任务
func (manage *EventManage) wake() {
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
func (manage *EventManage) AddTaskByExecuteList(_ []dto.Task) {
	manage.Lock()
	defer manage.Unlock()
}

// AddWaitExecuteEvent 添加任务到等待列表中
func (manage *EventManage) AddWaitExecuteEvent(events []model.Event) error {
	manage.Lock()
	defer manage.Unlock()
	for _, event := range events {
		logging.GLogger.Info(fmt.Sprintf("当前处理的任务信息:name:%s,id:%d,action:%d,status:%d", event.Name, event.ID, event.Action, event.Status))
		atomic.AddInt32(&manage.waitCount, 1)
		manage.EventStatusList.Store(event.ID, model.WAITING)
		manage.waitTask.LoadOrStore(event.ID, event)
	}
	return nil
}
func (manage *EventManage) Start() {
	go manage.Loop()
}

// GetTaskExecuteLave 获取任务空闲数
func (manage *EventManage) GetTaskExecuteLave() int32 {
	count := atomic.LoadInt32(&manage.executorCount)
	return manage.maxExecutorCount - count
}
