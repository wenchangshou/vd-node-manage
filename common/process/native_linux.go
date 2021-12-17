package process

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

func KillUE4(processID uint32) {

}

type UnixProcess struct {
	pid    int
	ppid   int
	state  rune
	pgrp   int
	sid    int
	binary string
}

func (p *UnixProcess) PPid() int {
	return p.ppid
}

func (p *UnixProcess) Pid() int {
	return p.pid
}
func (p *UnixProcess) Executable() string {
	return p.binary
}
func (p *UnixProcess) Refresh() error {
	statPath := fmt.Sprintf("/proc/%d/stat", p.pid)
	dataBytes, err := ioutil.ReadFile(statPath)
	if err != nil {
		return err
	}
	data := string(dataBytes)
	binStart := strings.IndexRune(data, '(') + 1
	binEnd := strings.IndexRune(data[binStart:], ')')
	p.binary = data[binStart : binStart+binEnd]
	data = data[binStart+binEnd+2:]
	_, err = fmt.Sscanf(data,
		"%c %d %d %d",
		&p.state,
		&p.ppid,
		&p.pgrp,
		&p.sid)
	return err
}
func findProcess(pid int) (Process, error) {
	dir := fmt.Sprintf("/proc/%d", pid)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return newUnixProcess(pid)
}
func newUnixProcess(pid int) (*UnixProcess, error) {
	p := &UnixProcess{pid: pid}
	return p, p.Refresh()
}
func GetUe4ProcessId(processID uint32) (int, error) {
	return -1, nil
}

func StartProcessAsCurrentUser(appPath, cmdLine, workDir string, backstage bool) (uint32, error) {
	cwd, _ := os.Getwd()
	re, _ := regexp.Compile("\\s{2,}")
	re2, _ := regexp.Compile("\"")
	cmdLine = re.ReplaceAllString(cmdLine, " ")
	cmdLine = re2.ReplaceAllString(cmdLine, "")
	arg := strings.Split(cmdLine, " ")
	cmd := exec.Command(appPath, arg...)
	cmd.Dir = cwd
	var out bytes.Buffer
	// cmd.Stdout = &out
	cmd.Stderr = &out
	// output, err := cmd.Output()
	err := cmd.Start()
	if err != nil {
		return 0, err
	}
	pid := cmd.Process.Pid
	// err = cmd.Process.Release()
	//if err!=nil{
	//	return 0,err
	//}
	go func() {
		cmd.Wait()
		pGid, err := syscall.Getpgid(cmd.Process.Pid)
		if err == nil {
			syscall.Kill(-pGid, 15)
		}
	}()
	return uint32(pid), nil
}
func CheckThreadExists(id uint32) bool {
	_, err := syscall.Getpgid(int(id))
	return err == nil
}
func KillProcesses(ps []int) {
	err := syscall.Kill(ps[0], 0)
	if err != nil {
		fmt.Println(fmt.Sprintf("kill process:%d,error:%s", ps[0], err.Error()))
	}
}
func KillPPT() {
}
func processes() ([]Process, error) {
	d, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer d.Close()
	results := make([]Process, 0, 50)
	for {
		names, err := d.Readdirnames(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, name := range names {
			if name[0] < '0' || name[0] > '9' {
				continue
			}
			pid, err := strconv.ParseInt(name, 10, 0)
			if err != nil {
				continue
			}
			p, err := newUnixProcess(int(pid))
			if err != nil {
				continue
			}
			results = append(results, p)
		}
	}
	return results, nil
}

func GetProcessIdByName(name string) (int, error) {
	processList, err := processes()
	if err != nil {
		return 0, nil
	}
	for x := range processList {
		var process Process = processList[x]
		if process.Executable() == name {
			return process.Pid(), nil
		}
	}
	return 0, errors.New("没有找到对应的进程")
}
