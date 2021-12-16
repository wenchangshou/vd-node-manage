package layout

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/model"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/pkg/e"
	"github.com/wenchangshou2/zutil"
	"sync"
)

type IManage interface {
	GetLayoutID() string
	OpenLayout(params model.OpenLayoutCmdParams) error
	CloseLayout() error
}
type layoutManage struct {
	sync.RWMutex
	layoutID string
	windows  map[string]*Window
}

// GetLayoutID 获取布局ID
func (manage *layoutManage) GetLayoutID() string {
	manage.RLock()
	defer manage.RUnlock()
	return manage.layoutID
}

// Kill 关闭已开启的程序
func (manage *layoutManage) Kill() {
	manage.Lock()
	defer manage.Unlock()
	for _, win := range manage.windows {
		win.Close()
		delete(manage.windows, manage.layoutID)
	}

}

// OpenLayout 打开布局
func (manage *layoutManage) OpenLayout(params model.OpenLayoutCmdParams) error {
	var (
		ports []int
		err   error
	)
	// 如果当前
	if params.ID != manage.layoutID {
		manage.Kill()
	}
	windowCount := len(params.Windows)
	if ports, err = zutil.GetFreePorts(windowCount); err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(windowCount)
	for k, w := range params.Windows {
		var (
			win *Window
			err error
		)
		if win, err = MakeWindow(w.ID, w.X, w.Y, w.Width, w.Height,
			w.Z, w.Service, w.Arguments, w.Source); err != nil {
			fmt.Println("open error")
		}
		win.Open(&wg, ports[k])
		manage.windows[win.ID] = win
	}
	manage.setLayout(params.ID)
	//for k, win := range windows {
	//	go func(win Window, port int) {
	//		fmt.Printf("open port:%d\n", port)
	//		win.Open(&wg, port)
	//	}(win, ports[k])
	//}
	wg.Wait()
	//manage.setLayout(id, windows)
	return nil
}
func (manage *layoutManage) setLayout(id string) {
	manage.Lock()
	defer manage.Unlock()
	manage.layoutID = id
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
		windows: make(map[string]*Window),
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
