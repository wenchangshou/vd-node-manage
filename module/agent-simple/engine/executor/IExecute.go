package executor

import (
	"github.com/wenchangshou2/vd-node-manage/common/model"
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
