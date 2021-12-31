package layout

import (
	"errors"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/model"
	model2 "github.com/wenchangshou/vd-node-manage/module/agent-simple/g/model"
	"github.com/wenchangshou2/zutil"
	"sync"
)

const (
	Run = iota
	Close
)

type WindowRunInfo struct {
	Run    bool   `json:"run"`
	Info   string `json:"info"`
	Source *Window
}

type ActiveLayoutInfo struct {
	ID      string `json:"active_id"`
	Status  bool   `json:"status"`
	Windows map[string]*model.ActiveWindowInfo
}

type WindowMap map[string]WindowRunInfo
type IManage interface {
	GetLayoutID() string
	GetLayoutRunInfo() ([]*model.ActiveWindowInfo, error)
	OpenLayout(params model.OpenLayoutCmdParams) error
	Control(params model.ControlWindowCmdParams, reply bool) error
	CloseLayout() error
}
type layoutManage struct {
	sync.RWMutex
	layoutID string
	windows  map[string]WindowRunInfo
	wSync    *sync.RWMutex
}

func (manage *layoutManage) GetLayoutRunInfo() (result []*model.ActiveWindowInfo, err error) {
	result = make([]*model.ActiveWindowInfo, 0)
	for wid, win := range manage.windows {
		var (
			info   string
			status bool
		)
		aWindow := model.ActiveWindowInfo{ID: wid}
		if info, err = win.Source.Get(); err != nil {
			continue
		}
		aWindow.Info = info
		if status, err = win.Source.GetRunStatus(); err != nil {
			continue
		}
		aWindow.Run = status
		result = append(result, &aWindow)
	}
	return result, nil
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
	for _, dst := range manage.windows {
		dst.Source.Close()
		delete(manage.windows, manage.layoutID)
	}

}
func (manage *layoutManage) Control(params model.ControlWindowCmdParams, reply bool) error {
	if manage.GetLayoutID() != params.ID {
		return errors.New("控制的非当前活动的布局")
	}
	win := manage.getWindow(params.Wid)
	if win == nil {
		return errors.New("未找到指定的窗口")
	}
	win.Source.Control(params.Body)
	return nil
}

func (manage *layoutManage) setWindow(wid string, win *Window) {
	manage.wSync.Lock()
	defer manage.wSync.Unlock()
	info := WindowRunInfo{Run: false, Source: win}
	manage.windows[wid] = info
}
func (manage *layoutManage) getWindow(wid string) *WindowRunInfo {
	manage.wSync.RLock()
	defer manage.wSync.RUnlock()
	w, ok := manage.windows[wid]
	if !ok {
		return nil
	}
	return &w
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
		manage.setWindow(win.ID, win)
	}
	manage.setLayout(params.ID)
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
		fmt.Println(window.Source.player.GetThreadId())
		window.Source.player.Close()
	}
	return nil
}

func InitLayoutManage() IManage {
	return &layoutManage{
		windows: make(map[string]WindowRunInfo),
		wSync:   new(sync.RWMutex),
	}
}

// Info 当前布局信息
type Info struct {
	// 布局id
	LayoutId string
	// 当前布局所打开的窗口
	Windows map[string]model2.Win
}
type ApplicationStatusChangeMsg struct {
	WinId     string
	ProcessId int
	Result    model2.PlayerArgumentInfo
	Type      string
}

func (w *WindowMap) Iterator(handle func(model2.Win) error) {
	//for _, win := range *w {
	//handle(win)
	//}

}
