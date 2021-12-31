package engine

import (
	"context"

	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/engine/executor"
)

type EventExecuteManage struct {
	event      model.Event
	ctx        context.Context
	statusChan chan model.EventStatus
	generator  executor.GeneratorFunction
}

func NewEventExecuteManage(event model.Event, ctx context.Context, exec executor.GeneratorFunction) *EventExecuteManage {
	t := &EventExecuteManage{
		event:      event,
		ctx:        ctx,
		statusChan: make(chan model.EventStatus),
		generator:  exec,
	}
	return t
}
func (task *EventExecuteManage) loop() {
	<-task.ctx.Done()
	task.statusChan <- executor.CANCEL
}
func (task *EventExecuteManage) execute() {
	e := task.event
	execFunc, err := task.generator(e)
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
func (task *EventExecuteManage) Start() chan model.EventStatus {
	c := make(chan model.EventStatus)
	task.statusChan = c
	go task.loop()
	go task.execute()
	return c
}
