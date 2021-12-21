package player

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/player/playerService"
	"path"
	"sync"

	"github.com/wenchangshou2/vd-node-manage/common/process"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	"github.com/wenchangshou2/zutil"
)

type ResourcePlayer struct {
	e.Window
	Source    string
	Arguments map[string]interface{}
	Pid       int
	port      int
	PlayPath  string
	service   playerService.IPlayerService
}

func (player *ResourcePlayer) Control(body string) (string, error) {
	if player.service == nil {
		return "", nil
	}
	return player.service.Control(body)
}

func (player *ResourcePlayer) GetThreadId() int {
	return player.Pid
}
func (_ *ResourcePlayer) Check() (bool, error) {
	return false, nil
}

// Open 打开一个播放器
func (player *ResourcePlayer) Open(wg *sync.WaitGroup, p int) (pid int, err error) {
	e := process.StandardApplicationControl{}
	defer wg.Done()
	params := ""
	if player.Arguments != nil && len(player.Arguments) > 0 {
		params = zutil.MapToString(player.Arguments)
	}
	source := path.Join(g.Config().Resource.Directory, "resource", player.Source)
	params = fmt.Sprintf("%s -w %d -h %d -x %d -y %d", params, player.Width, player.Height, player.X, player.Y)
	params = fmt.Sprintf("%s -source %s", params, source)
	params += fmt.Sprintf(" -http true -httpPort %d", p)
	player.port = p
	fmt.Println(player.PlayPath, params)
	player.Pid, err = e.StartProcessAsCurrentUser(player.PlayPath, params, "", false)
	service, err := playerService.GeneratePlayerService("http", p)
	if err != nil {
		return 0, err
	}
	player.service = service
	return player.Pid, err
}

// Close 关闭播放器
func (player *ResourcePlayer) Close() error {
	fmt.Println("pid", player.Pid)
	if player.Pid == 0 {
		return nil
	}
	process.KillProcesses([]int{int(player.Pid)})
	return nil
}
