package engine

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/executor"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
)

type eventExecuteManage struct {
	event      model.Event
	ctx        context.Context
	statusChan chan model.EventStatus
	generator  executor.GeneratorFunction
}

func NewEventExecuteManage(event model.Event, ctx context.Context, exec executor.GeneratorFunction) *eventExecuteManage {
	t := &eventExecuteManage{
		event:      event,
		ctx:        ctx,
		statusChan: make(chan model.EventStatus),
		generator:  exec,
	}
	return t
}
func (task *eventExecuteManage) loop() {
	select {
	case <-task.ctx.Done():
		task.statusChan <- executor.CANCEL
	}

}
func (task eventExecuteManage) action(t e.TaskItem) {

}
func (task *eventExecuteManage) execute() {
	//l := NewTaskLinedList()
	//for _, item := range task.src.Items {
	//	// 如果
	//	if item.Depend == "" {
	//		l.Append(&DoubleNode{Data: DoubleObject(item)})
	//		continue
	//	}
	//	idx, exists := l.GetByTaskId(item.Depend)
	//	if !exists {
	//		l.Insert(0, &DoubleNode{Data: DoubleObject(item)})
	//		continue
	//	}
	//	l.Insert(idx, &DoubleNode{Data: DoubleObject(item)})
	//}
	//l.Foreach(func(node *DoubleNode) bool {
	//	item := node.Data
	//	e, err := task.generator(e.ExecuteType(item.Action), item.ID, item.Options)
	//	if err != nil {
	//		logging.GLogger.Warn(fmt.Sprintf("生成类别错误:%s", err.Error()))
	//		task.statusChan <- executor.ERROR
	//		return false
	//	}
	//	err = e.Execute()
	//	if err != nil {
	//		task.statusChan <- executor.ERROR
	//
	//		return false
	//	}
	//	return true
	//})
	//
	//task.statusChan <- executor.DONE
	e := task.event
	execFunc, err := task.generator(e.Action, e.ID, e.Params)
	if err != nil {
		task.statusChan <- model.Error
		return
	}
	if err = execFunc.Execute(); err != nil {
		task.statusChan <- model.Error
		return
	}
	task.statusChan <- model.Done
}

// Start 启动一个任务组
func (task *eventExecuteManage) Start() chan model.EventStatus {
	c := make(chan model.EventStatus)
	task.statusChan = c
	go task.loop()
	go task.execute()
	return c
}
