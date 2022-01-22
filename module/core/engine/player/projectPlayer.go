package player

import (
	"fmt"
	"github.com/wenchangshou/vd-node-manage/module/core/g/model"
	"github.com/wenchangshou/vd-node-manage/module/core/g/process"
	"sync"
)

type ProjectPlayer struct {
	model.Window
	Source  string
	end     chan bool
	Pid     uint32
	service string
}

func (player *ProjectPlayer) Open(wg *sync.WaitGroup, p int) (pid uint32, err error) {
	//_source := path.Join(g.Config().Resource.Directory, "player", player.Source)
	defer wg.Done()
	params := ""
	params = fmt.Sprintf(" %s -w %d -h %d -x %d -y %d", params, player.Width, player.Height, player.X, player.Y)
	player.Pid, err = process.GProcess.StartProcessAsCurrentUser(player.Source, params, "", false)
	return player.Pid, err
}

func (player ProjectPlayer) GetThreadId() uint32 {
	//TODO implement me
	return player.Pid
}

func (player ProjectPlayer) Close() error {
	if player.Pid == 0 {
		return nil
	}
	if player.service == "ue4" {
		process.KillUE4(player.Pid)
	} else {
		process.KillProcesses([]int{int(player.Pid)})
	}
	go func() {
		player.end <- true
	}()
	return nil
}

func (player ProjectPlayer) Check() (bool, error) {
	exists := process.GProcess.GetThreadStatus(player.Pid)
	return exists, nil
}

func (player ProjectPlayer) Control(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (player ProjectPlayer) Get() (string, error) {
	return "", nil
}

func (player ProjectPlayer) Change(s string) error {
	return nil
}
