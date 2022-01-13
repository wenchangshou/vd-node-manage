package process

import (
	"fmt"
	"os/exec"
	"strings"
)

type WindowConsoleProcess struct {
}

func (w WindowConsoleProcess) GetThreadStatus(id uint32) bool {
	return CheckThreadExists(id)
}

func (w WindowConsoleProcess) StartProcessAsCurrentUser(appPath, cmdLine, workDir string, backstage bool) (uint32, error) {
	var (
		result []byte
		err    error
	)
	params := strings.Split(cmdLine, " ")
	cmd := exec.Command(appPath, params...)
	cmd.Dir = workDir
	if err = cmd.Start(); err != nil {
		return 0, err
	}
	fmt.Println(string(result))
	return uint32(cmd.Process.Pid), nil
}

func NewWindowConsoleProcess() *WindowConsoleProcess {
	return &WindowConsoleProcess{}
}
