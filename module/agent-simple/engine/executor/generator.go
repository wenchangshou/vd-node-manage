package executor

import (
	"errors"
	"github.com/wenchangshou/vd-node-manage/common/model"
	IService "github.com/wenchangshou/vd-node-manage/module/agent-simple/service"
)

// GenerateExecutorFactoryFunc 生成执行器工厂函数
func GenerateExecutorFactoryFunc(serviceFactory *IService.ServiceFactory, httpRequestUri string) GeneratorFunction {
	return func(event model.Event) (IExecute, error) {
		//var err error
		switch event.Action {
		case model.InstallResourceAction:
			resource, err := serviceFactory.Resource.QueryResource(event.ResourceId)
			if err != nil {
				return nil, errors.New("rpc请求资源接口失败")
			}
			e := &InstallResourceExecutor{
				taskID:          event.ID,
				eventService:    serviceFactory.Event,
				DeviceService:   serviceFactory.Device,
				HttpRequestUri:  httpRequestUri,
				ResourceService: serviceFactory.Resource,
				Resource:        resource,
			}
			return e, nil
		case model.DeleteResource:
			resource, err := serviceFactory.Resource.QueryResource(event.ResourceId)
			if err != nil {
				return nil, errors.New("rpc请求资源接口失败")
			}
			e := &DeleteResourceExecutor{
				Resource:        resource,
				DeviceService:   serviceFactory.Device,
				eventService:    serviceFactory.Event,
				ResourceService: serviceFactory.Resource,
			}
			return e, nil

		default:
			return nil, errors.New("没有找到对应的执行程序")
		}
	}
}
