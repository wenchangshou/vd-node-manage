package layout

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	"github.com/wenchangshou2/zutil"
	"sync"
)

type IManage interface {
	GetLayoutID() string
	OpenLayout(string, []Window) error
	CloseLayout() error
}
type layoutManage struct {
	sync.RWMutex
	layoutID string
	windows  []Window
}

// GetLayoutID 获取布局ID
func (manage *layoutManage) GetLayoutID() string {
	manage.RLock()
	defer manage.RUnlock()
	return manage.layoutID
}

// OpenLayout 打开布局
func (manage *layoutManage) OpenLayout(id string, windows []Window) error {

	windowCount := len(windows)
	var wg sync.WaitGroup
	wg.Add(windowCount)
	ports, err := zutil.GetFreePorts(windowCount)
	if err != nil {
		return errors.Wrap(err, "获取空间端口失败")
	}
	for k, win := range windows {
		go func(win Window, port int) {
			fmt.Printf("open port:%d\n", port)
			win.Open(&wg, port)
		}(win, ports[k])
	}
	wg.Wait()
	manage.setLayout(id, windows)
	return nil
}
func (manage *layoutManage) setLayout(id string, wins []Window) {
	manage.Lock()
	defer manage.Unlock()
	manage.windows = wins
}
func (manage *layoutManage) CloseLayout() error {
	for _, window := range manage.windows {
		fmt.Println(window.player.GetThreadId())
		window.player.Close()
	}
	return nil
}

func InitLayoutManage() IManage {
	return &layoutManage{
		windows: make([]Window, 0),
	}
}

// Info 当前布局信息
type Info struct {
	// 布局id
	LayoutId string
	// 当前布局所打开的窗口
	Windows map[string]e.Win
}
type ApplicationStatusChangeMsg struct {
	WinId     string
	ProcessId int
	Result    e.PlayerArgumentInfo
	Type      string
}
type WindowMap map[string]e.Win

func (w *WindowMap) Iterator(handle func(e.Win) error) {
	for _, win := range *w {
		handle(win)
	}
}
