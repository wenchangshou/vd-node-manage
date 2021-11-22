package layout

import (
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine/player"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
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
	Arguments map[string]string
	Source    string
	player    player.IPlayer
}

func (window *Window) Open(wg *sync.WaitGroup, port int) error {
	return window.player.Open(wg, port)
}

func MakeWindow(id string, x int, y int, width int, height int, z int,
	service string,
	Arguments map[string]string,
	source string,
) (*Window, error) {
	windowInfo := e.Window{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Z:      z,
	}
	_player, err := player.MakePlayer(windowInfo, "", service, source)
	if err != nil {
		return nil, err
	}
	return &Window{
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