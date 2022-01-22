package process

import (
	"os"
	"strings"
)

type IProcess interface {
	StartProcessAsCurrentUser(appPath, cmdLine, workDir string, backstage bool) (uint32, error)
	GetThreadStatus(id uint32) bool
	KillUe4(pid uint32)
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
func KillUE4(processID uint32) {
	tree, _ := SnapshotSysProcesses()
	processName := ""
	parentId := -1
	for _, ps := range tree {
		if ps.ProcessID == int(processID) {
			processName = ps.ProcessName
			parentId = ps.ProcessID
		}
	}
	if len(processName) == 0 {
		return
	}
	for _, ps := range tree {
		if strings.Compare(processName, ps.ProcessName) == 0 || ps.ParentProcessID == parentId {
			KillProcesses([]int{ps.ProcessID})
		}
	}
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
