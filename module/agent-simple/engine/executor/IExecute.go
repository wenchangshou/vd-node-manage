package executor

import (
	"errors"
	"github.com/wenchangshou2/vd-node-manage/common/model"
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

type GeneratorFunction func(model.EventAction, uint, interface{}) (IExecute, error)

// IExecute 执行器接口
type IExecute interface {
	Execute() error
	BindOption(interface{}) error
	Verification(string) bool
}

// GenerateExecutorFactoryFunc 生成执行器工厂函数
func GenerateExecutorFactoryFunc(eventService IService.EventService, httpRequestUri string) GeneratorFunction {
	return func(executorType model.EventAction, taskID uint, option interface{}) (IExecute, error) {
		//var err error
		switch executorType {
		//case model.InstallProjectAction:
		//	e := &InstallProjectExecutor{
		//		TaskID:         taskID,
		//		HttpRequestUri: httpRequestUri,
		//		TaskService:    eventService,
		//	}
		//	if err = e.BindOption(option); err != nil {
		//		return nil, errors.New("install project action bind param error")
		//	}
		//	return e, nil
		//
		//case model.DeleteProject:
		//	e := &DeleteProjectExecutor{
		//		TaskID: taskID,
		//	}
		//	if err = e.BindOption(option); err != nil {
		//		return nil, errors.New("delete project action bind param error")
		//	}
		//	return e, nil

		case model.InstallResourceAction:
			e := &InstallResourceExecutor{
				taskID:         taskID,
				eventService:   eventService,
				HttpRequestUri: httpRequestUri,
			}
			e.BindOption(option)
			//if err = e.BindOption(option); err != nil {
			//	return nil, errors.New("install resource action bind param error")
			//}
			return e, nil
		case model.DeleteResource:
			e := &DeleteResourceExecutor{}
			//if err = e.BindOption(option); err != nil {
			//	return nil, errors.New("delete resource action bind param error")
			//}
			return e, nil

		default:
			return nil, errors.New("没有找到对应的执行程序")
		}
	}
}
