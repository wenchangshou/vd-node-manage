package executor

import (
	"errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	IService "github.com/wenchangshou2/vd-node-manage/module/agent-simple/service"
)

const (
	INITIALIZE = iota
	EXECUTE
	DONE
	ERROR
	CANCEL
	ALL
)

type GeneratorFunction func(e.ExecuteType, uint, string) (IExecute, error)

// IExecute 执行器接口
type IExecute interface {
	Execute() error
	BindOption(string) error
	Verification(string) bool
}

// GenerateExecutorFactoryFunc 生成执行器工厂函数
func GenerateExecutorFactoryFunc(taskService IService.TaskService, httpRequestUri string) GeneratorFunction {
	return func(executorType e.ExecuteType, taskID uint, option string) (IExecute, error) {
		var err error
		switch executorType {
		case e.InstallProjectAction:
			e := &InstallProjectExecutor{
				TaskID:         taskID,
				HttpRequestUri: httpRequestUri,
				TaskService:    taskService,
			}
			if err = e.BindOption(option); err != nil {
				return nil, errors.New("install project action bind param error")
			}
			return e, nil

		case e.DeleteProject:
			e := &DeleteProjectExecutor{
				TaskID: taskID,
			}
			if err = e.BindOption(option); err != nil {
				return nil, errors.New("delete project action bind param error")
			}
			return e, nil
		case e.InstallResourceAction:
			e := &InstallResourceExecutor{
				taskID:         taskID,
				taskService:    taskService,
				HttpRequestUri: httpRequestUri,
			}
			if err = e.BindOption(option); err != nil {
				return nil, errors.New("install resource action bind param error")
			}
			return e, nil
		case e.DeleteResource:
			e := &DeleteResourceExecutor{}
			if err = e.BindOption(option); err != nil {
				return nil, errors.New("delete resource action bind param error")
			}
			return e, nil

		default:
			return nil, errors.New("没有找到对应的执行程序")
		}
	}
}
