package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/getlantern/systray"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/cron"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/engine"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/g"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/gui/icon"
	"github.com/wenchangshou2/vd-node-manage/module/agent-simple/http"
	"github.com/wenchangshou2/zutil"
)

var (
	confPath      string
	installFlag   bool
	uninstallFlag bool
)

// waitRegister 等待注册
func waitRegister() <-chan bool {
	result := make(chan bool)
	go func() {
		timeTicker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-timeTicker.C:
				if g.Config().Server.Register {
					result <- true
				}
			}
		}
	}()

	return result
}
func InitSystemInfo(cfg *string, hardware *string) {
	var (
		err error
	)
	g.InitApplication()
	g.ParseConfig(*cfg)
	conf := g.Config()
	logging.InitLogging(conf.Log.Path, conf.Log.Level)
	// 创建工作目录
	ap := path.Join(conf.Resource.Directory, "application")
	rp := path.Join(conf.Resource.Directory, "resource")
	zutil.IsNotExistMkDir(conf.Resource.Directory)
	zutil.IsNotExistMkDir(conf.Resource.Tmp)
	zutil.IsNotExistMkDir(ap)
	zutil.IsNotExistMkDir(rp)
	g.ParseHardware(*hardware)
	go http.Start()
	<-waitRegister()
	fmt.Println("注册成功")
	g.InitLocalIp()
	g.InitRpcClients()
	cron.ReportDeviceStatus()
	if err = engine.InitSchedule(conf); err != nil {
		log.Fatalln("初始化调度失败")
	}
}
func InitUi() {
	a := app.New()
	w := a.NewWindow("Hello")
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))
	fmt.Println("end")

}
func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("vd 播控系统")
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
	}()

}

func main() {

	cfg := flag.String("c", zutil.RelativePath("cfg.json"), "configuration file")
	hardware := flag.String("d", zutil.RelativePath("hardware.data"), "hardware file")
	flag.BoolVar(&installFlag, "i", false, "install program")
	flag.BoolVar(&uninstallFlag, "u", false, "uninstall program")
	flag.Parse()
	InitSystemInfo(cfg, hardware)
	onExit := func() {
		os.Exit(0)
	}
	go systray.Run(onReady, onExit)
	fmt.Println("icon init")
	InitUi()
	select {}
	// Init
}
