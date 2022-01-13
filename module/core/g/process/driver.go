package process

import "os"

type IProcess interface {
	StartProcessAsCurrentUser(appPath, cmdLine, workDir string, backstage bool) (uint32, error)
	GetThreadStatus(id uint32) bool
}

var GProcess IProcess

func GenerateProcess(driver string) error {
	if driver == "console" {
		GProcess = NewWindowConsoleProcess()
	} else if driver == "service" {
		GProcess = NewWindowServiceProcess()
	}
	return nil
}

// KillProcesses 杀死单个进程
func KillProcesses(ps []int) {
	for _, pid := range ps {
		p, err := os.FindProcess(pid)
		if err != nil {
			continue
		}
		p.Kill()
		p.Release()
	}
}
