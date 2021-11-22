package engine

import (
	"context"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent/dto"
	"github.com/wenchangshou2/vd-node-manage/module/agent/engine/executor"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
)

type TaskGroup struct {
	src        dto.Task
	ctx        context.Context
	statusChan chan int
	generator  executor.GeneratorFunction
}

func NewTaskGroup(task dto.Task, ctx context.Context, exec *executor.GeneratorFunction) *TaskGroup {
	t := &TaskGroup{
		src:        task,
		ctx:        ctx,
		statusChan: make(chan int),
		generator:  *exec,
	}
	return t
}
func (task *TaskGroup) loop() {
	select {
	case <-task.ctx.Done():
		task.statusChan <- executor.CANCEL
	}

}
func (task TaskGroup) action(t e.TaskItem) {

}
func (task *TaskGroup) execute() {
	l := NewTaskLinedList()
	for _, item := range task.src.Items {
		// 如果
		if item.Depend == "" {
			l.Append(&DoubleNode{Data: DoubleObject(item)})
			continue
		}
		idx, exists := l.GetByTaskId(item.Depend)
		if !exists {
			l.Insert(0, &DoubleNode{Data: DoubleObject(item)})
			continue
		}
		l.Insert(idx, &DoubleNode{Data: DoubleObject(item)})
	}
	l.Foreach(func(node *DoubleNode) bool {
		item := node.Data
		e, err := task.generator(e.ExecuteType(item.Action), item.ID, item.Options)
		if err != nil {
			logging.GLogger.Warn(fmt.Sprintf("生成类别错误:%s", err.Error()))
			task.statusChan <- executor.ERROR
			return false
		}
		err = e.Execute()
		if err != nil {
			task.statusChan <- executor.ERROR

			return false
		}
		return true
	})
	//for _, item := range task.src.Items {
	//	e, err := task.generator(e.ExecuteType(item.Action), item.ID, item.Options)
	//	if err != nil {
	//		logging.G_Logger.Warn(fmt.Sprintf("生成类别错误:%s", err.Error()))
	//		task.statusChan <- executor.ERROR
	//		return
	//	}
	//	err = e.Execute()
	//	if err != nil {
	//		task.statusChan <- executor.ERROR
	//		return
	//	}
	//	fmt.Println(err)
	//}
	task.statusChan <- executor.DONE
}

// Start 启动一个任务组
func (task *TaskGroup) Start() chan int {
	c := make(chan int)
	task.statusChan = c
	go task.loop()
	go task.execute()
	return c
}
