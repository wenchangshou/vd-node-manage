package layout

import (
	"github.com/wenchangshou/vd-node-manage/module/core/engine/player"
	"github.com/wenchangshou/vd-node-manage/module/core/g"
	"github.com/wenchangshou/vd-node-manage/module/core/g/model"
	"path"
	"sync"
)

type WindowStatus int

func (w WindowStatus) String() string {
	return [...]string{"Init", "Running", "Opening", "Close", "Abnormal"}[w]
}

type WindowStyle struct {
	WindowStyle string `json:"WindowStyle"`
}
type Window struct {
	ID        string
	X         int
	Y         int
	Width     int
	Height    int
	Z         int
	Service   string
	Status    WindowStatus
	Style     WindowStyle
	Arguments map[string]interface{}
	Source    string
	player    player.IPlayer
	port      int
	win       model.Window
}

func (window *Window) Open(wg *sync.WaitGroup, port int) (uint32, error) {
	window.port = port
	return window.player.Open(wg, port)
}
func (window *Window) Close() error {
	return window.player.Close()
}
func (window *Window) Control(body string) (string, error) {
	return window.player.Control(body)
}
func (window *Window) Change(source string) error {
	return window.player.Change(source)
}
func (window *Window) ChangePlayer(service, source string) (pid uint32, err error) {
	window.player.Close()
	var wg sync.WaitGroup
	_source := ""
	if service == "web" {
		_source = source
	} else if service == "app" || service == "ue4" {
		_source = path.Join(g.Config().Resource.Directory, "project", source)
	} else {
		_source = path.Join(g.Config().Resource.Directory, "resource", source)
	}
	if window.player, err = player.MakePlayer(window.win, "", service, _source); err != nil {
		return
	}
	wg.Add(1)
	pid, err = window.Open(&wg, window.port)
	window.Service = service
	wg.Wait()
	return
}
func (window *Window) Get() (string, error) {
	return window.player.Get()
}
func (window *Window) GetRunStatus() (bool, error) {
	return window.player.Check()
}
func MakeWindow(id string, x int, y int, width int, height int, z int,
	service string,
	Arguments map[string]interface{},
	source string,
) (*Window, error) {
	var (
		_player player.IPlayer
		err     error
		_source string
	)
	windowInfo := model.Window{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Z:      z,
	}
	if service == "web" {
		_source = source
	} else if service == "app" || service == "ue4" {
		_source = path.Join(g.Config().Resource.Directory, "project", source)
	} else {
		_source = path.Join(g.Config().Resource.Directory, "resource", source)
	}
	if _player, err = player.MakePlayer(windowInfo, "", service, _source); err != nil {
		return nil, err
	}
	return &Window{
		ID:        id,
		win:       windowInfo,
		Source:    source,
		Service:   service,
		Arguments: Arguments,
		player:    _player,
	}, nil
}
