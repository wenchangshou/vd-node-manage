package executor

import "github.com/wenchangshou/vd-node-manage/common/model"

const (
	INITIALIZE = iota
	EXECUTE
	DONE
	ERROR
	CANCEL
	ALL
)

type GeneratorFunction func(event model.Event) (IExecute, error)

// IExecute 执行器接口
type IExecute interface {
	Cancel() error
	Execute() error
	BindOption(interface{}) error
	Verification(string) bool
}
