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
	<-task.ctx.Done()

}
func (task eventExecuteManage) action(_ e.TaskItem) {

}
func (task *eventExecuteManage) execute() {
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
