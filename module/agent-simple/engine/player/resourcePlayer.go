package player

import (
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/logging"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/engine/player/playerService"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g/process"
	"path"
	"sync"
	"time"

	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/zutil"
)

type ResourcePlayer struct {
	model.Window
	Source    string
	Arguments map[string]interface{}
	Pid       uint32
	port      int
	PlayPath  string
	service   playerService.IPlayerService
	end       chan bool
	info      string
}

func (player *ResourcePlayer) Get() (string, error) {
	return player.service.Get()
}

func (player *ResourcePlayer) Control(body string) (string, error) {
	if player.service == nil {
		return "", nil
	}
	return player.service.Control(body)
}
func (player *ResourcePlayer) Loop() {
	d := time.NewTicker(time.Second)
	for {
		select {
		case <-d.C:
			if info, err := player.service.Get(); err == nil {
				player.info = info
			}
		case <-player.end:
			return
		}
		time.Sleep(time.Second)
	}
}
func (player *ResourcePlayer) GetThreadId() uint32 {
	return player.Pid
}
func (_ *ResourcePlayer) Check() (bool, error) {
	return false, nil
}

// Open 打开一个播放器
func (player *ResourcePlayer) Open(wg *sync.WaitGroup, p int) (pid uint32, err error) {
	var (
		service playerService.IPlayerService
	)
	defer wg.Done()
	params := ""
	if player.Arguments != nil && len(player.Arguments) > 0 {
		params = zutil.MapToString(player.Arguments)
	}
	source := path.Join(g.Config().Resource.Directory, "resource", player.Source)
	params = fmt.Sprintf(" %s -w %d -h %d -x %d -y %d", params, player.Width, player.Height, player.X, player.Y)
	params += fmt.Sprintf(" -s %s", source)
	//params += fmt.Sprintf(" - true -httpPort %d", p)
	params += fmt.Sprintf(" -p %d", p)
	player.port = p
	logging.GLogger.Info(fmt.Sprintf("player path:%s,arguments:%s", player.PlayPath, params))
	player.Pid, err = process.GProcess.StartProcessAsCurrentUser(player.PlayPath, params, "", false)
	if service, err = playerService.GeneratePlayerService("rpc", p); err != nil {
		return 0, err
	}
	player.service = service
	go player.Loop()
	return player.Pid, err
}

// Close 关闭播放器
func (player *ResourcePlayer) Close() error {
	fmt.Println("pid", player.Pid)
	if player.Pid == 0 {
		return nil
	}
	process.KillProcesses([]int{int(player.Pid)})
	go func() {
		player.end <- true
	}()
	return nil
}
