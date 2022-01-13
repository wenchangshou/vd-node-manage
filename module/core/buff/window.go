package buff

import (
	"sync"
	"time"
)

type WindowGlobalStatus struct {
	lastUpTime    map[string]int64
	windowRunInfo map[string]string
	sync.RWMutex
}

func (g *WindowGlobalStatus) SetWindowHealth(wid string) {
	msec := time.Now().UnixNano() / 1000000
	g.Lock()
	defer g.Unlock()
	g.lastUpTime[wid] = msec
}
func (g *WindowGlobalStatus) SetWindowStatus(wid string, info string) {
	g.Lock()
	defer g.Unlock()
	g.windowRunInfo[wid] = info
}

var (
	GWindowGlobalStatus *WindowGlobalStatus
)

func InitGlobalBuffer() {
	GWindowGlobalStatus = &WindowGlobalStatus{}
}
