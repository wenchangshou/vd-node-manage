package playerService

import "errors"

type IPlayerService interface {
	Ping() (bool, error)
	Control(string) (string, error)
	Get() (string, error)
}

func GeneratePlayerService(service string, port int) (IPlayerService, error) {
	if service == "http" {
		return &HttpPlayerService{port}, nil
	} else if service == "rpc" {
		return &RpcPlayerService{port: port}, nil
	}
	return nil, errors.New("未知服务")
}
