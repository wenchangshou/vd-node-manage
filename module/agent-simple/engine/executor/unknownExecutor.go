package executor

type UnknownExecutor struct {
}

func (executor *UnknownExecutor) Execute() error {
	return nil
}
func (executor *UnknownExecutor) Cancel() error {
	return nil
}
func (executor *UnknownExecutor) Verification(_ string) bool {
	return false
}
func (executor *UnknownExecutor) SubscribeNotifyStatusChange(func(string, int, string)) {

}

// BindOption  检验任务参数
func (executor *UnknownExecutor) BindOption(_ string) error {
	return nil
}
