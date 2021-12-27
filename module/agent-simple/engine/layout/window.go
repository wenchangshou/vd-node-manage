package layout

import (
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/player"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g/model"
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
}

func (window *Window) Open(wg *sync.WaitGroup, port int) (int, error) {
	return window.player.Open(wg, port)
}
func (window *Window) Close() error {
	return window.player.Close()
}
func (window *Window) Control(body string) (string, error) {
	return window.player.Control(body)
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
	)
	windowInfo := model.Window{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Z:      z,
	}
	if _player, err = player.MakePlayer(windowInfo, "", service, source); err != nil {
		return nil, err
	}
	return &Window{
		ID:        id,
		X:         x,
		Y:         y,
		Width:     width,
		Height:    height,
		Z:         z,
		Source:    source,
		Arguments: Arguments,
		player:    _player,
	}, nil
}
