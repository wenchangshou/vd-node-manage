package process

import (
	"fmt"
	"os/exec"
	"strings"
)

type StandardApplicationControl struct {
}

func (p StandardApplicationControl) StartProcessAsCurrentUser(appPath, cmdLine, workDir string, _ bool) (int, error) {
	var (
		result []byte
		err    error
	)
	cmds := strings.Split(cmdLine, " ")
	fmt.Println("exec", appPath, cmds)
	cmd := exec.Command(appPath, cmds...)
	cmd.Dir = workDir
	if err = cmd.Start(); err != nil {
		return 0, err
	}
	fmt.Println(string(result))
	return cmd.Process.Pid, nil
}
