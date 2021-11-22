package player

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/process"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
	"github.com/wenchangshou2/zutil"
	"sync"
)

type ResourcePlayer struct {
	e.Window
	Source    string
	Arguments map[string]interface{}
	Pid       uint32
	PlayPath  string
}

func (player *ResourcePlayer) GetThreadId() uint32 {
	return player.Pid
}
func (player *ResourcePlayer) Check() (bool, error) {
	return false, nil
}

// Open 打开一个播放器
func (player *ResourcePlayer) Open(wg *sync.WaitGroup, port int) (err error) {
	fmt.Println("open")
	defer wg.Done()
	params := ""
	if player.Arguments != nil && len(player.Arguments) > 0 {
		params = zutil.MapToString(player.Arguments)
	}
	params = fmt.Sprintf("-w %d -h %d -x %d -y %d", player.Width, player.Height, player.X, player.Y)
	params = fmt.Sprintf("%s -source %s", params, player.Source)
	fmt.Println(player.PlayPath, params)
	player.Pid, err = process.StartProcessAsCurrentUser(player.PlayPath, params, "", false)
	return
}

func (player *ResourcePlayer) Close() error {
	fmt.Println("pid", player.Pid)
	if player.Pid == 0 {
		return nil
	}
	process.KillProcesses([]int{int(player.Pid)})
	return nil
}
