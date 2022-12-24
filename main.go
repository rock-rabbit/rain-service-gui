package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"syscall"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rock-rabbit/rain-service-gui/module"
)

var (
	ErrorDialog func(err error)
	Rain        = NewRainService(config.RainService)
	RainCmd     *exec.Cmd

	// App 服务
	App fyne.App

	// 主窗口
	MasterW fyne.Window

	// 设置窗口
	Setting    = false
	SettingW   fyne.Window
	SettingBtn *widget.Button

	Error error
)

func main() {
	var err error
	// 初始化日志
	os.MkdirAll(GetLogsDir(), os.ModePerm)
	f, err := os.OpenFile(filepath.Join(GetLogsDir(), "rain.log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// 启动服务
	go func() {
		exePath := GetExecutable()
		log.Println(exePath)

		service := "rain-service.exe"
		if runtime.GOOS == "darwin" {
			service = "rain-service"
		}
		servicePath := filepath.Join(exePath, service)
		log.Println(servicePath)
		if FileNotExist(servicePath) {
			log.Println("不存在本地 rain-service")
			return
		}
		_, err := Rain.GetRow("")
		if err == nil {
			log.Println("已经运行了 rain-service 服务")
			return
		}
		// 启动
		log.Println("启动 rain-service 服务")
		RainCmd = exec.Command(servicePath)
		err = RainCmd.Start()
		if err != nil {
			log.Println("RainCmd.Start", err)
		}
		log.Println("启动 rain-service 服务成功")
	}()

	// 窗口
	App = app.New()
	App.Settings().SetTheme(&myTheme{})

	MasterW = App.NewWindow("Rain")
	MasterW.SetContent(makeUI())
	MasterW.SetFixedSize(true)
	MasterW.SetMaster()
	MasterW.SetOnClosed(func() {
		// 退出 APP
		if RainCmd != nil {
			log.Println("退出 rain-service 服务")
			err = RainCmd.Process.Signal(syscall.SIGINT)
			if err != nil {
				log.Println("RainCmd.Process.Signal", err)
			}
		}
	})

	// 错误提示框
	ErrorDialog = func(err error) {
		dialog.ShowError(err, MasterW)
	}
	MasterW.ShowAndRun()
}

// makeSetting 设置
func makeSetting() {
	// 仅支持开启一个设置窗口
	if Setting {
		return
	}

	// 已开启
	Setting = true
	SettingBtn.Disable()

	SettingW = App.NewWindow("设置")
	SettingW.Resize(fyne.NewSize(600, 300))
	SettingW.SetFixedSize(true)
	SettingW.Show()

	form := container.New(layout.NewFormLayout())

	// 下载输出目录
	downloadDirE := widget.NewEntry()
	downloadDirE.SetText(config.Outdir)
	form.Add(widget.NewLabel("下载目录"))
	form.Add(downloadDirE)

	// 服务器地址
	rainHostE := widget.NewEntry()
	rainHostE.SetText(config.RainService)
	form.Add(widget.NewLabel("服务地址"))
	form.Add(rainHostE)

	// 刷新时间
	refreshTimeE := widget.NewEntry()
	refreshTimeE.SetText(fmt.Sprintf("%d", config.RefreshTime))
	form.Add(widget.NewLabel("刷新时间"))
	form.Add(refreshTimeE)

	form.Add(widget.NewLabel(""))
	form.Add(widget.NewButton("应用", func() {
		err := SaveConfig(
			downloadDirE.Text,
			refreshTimeE.Text,
			rainHostE.Text,
		)
		if err != nil {
			dialog.ShowError(err, SettingW)
			return
		}
		dialog.ShowInformation("提示", "应用成功 :) ", SettingW)
	}))
	SettingW.SetContent(form)

	// 关闭设置窗口后
	SettingW.SetOnClosed(func() {
		Setting = false
		SettingBtn.Enable()
	})
}

// makeUI 创建内容
func makeUI() fyne.CanvasObject {
	vbox := container.NewVBox()

	// 标题
	SettingBtn = widget.NewButtonWithIcon("设置", theme.SettingsIcon(), makeSetting)
	hbox := container.NewHBox(
		widget.NewLabel("Rain 下载器"),
		layout.NewSpacer(),
		SettingBtn,
	)
	vbox.Add(hbox)

	// 工具栏
	vbox.Add(makeTools())

	// 下载列表
	dl := module.NewDownloadList()
	vbox.Add(widget.NewLabel("下载列表"))
	vbox.Add(dl)
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(config.RefreshTime)) {
			rows, _ := Rain.GetRow("")
			dl_rows := make([]*module.Row, 0, len(rows))
			for _, v := range rows {
				uuid := v.UUID
				dl_rows = append(dl_rows, &module.Row{
					UUID:            uuid,
					Title:           GetFilename(v.Stat.Outpath),
					Progess:         float64(v.Stat.Progress),
					Status:          v.Status,
					TotalLength:     v.Stat.TotalLength,
					CompletedLength: v.Stat.CompletedLength,
					DownloadSpeed:   v.Stat.DownloadSpeed,
					Click: func(c string) {
						switch c {
						case "暂停":
							Rain.Pause(uuid)
						case "继续", "重试", "重下":
							Rain.Start(uuid)
						case "删除":
							dialog.ShowConfirm(
								"信息",
								"是否要删除 "+GetFilename(v.Stat.Outpath)+" ？",
								func(b bool) {
									if b {
										Rain.Delete(uuid)
									}
								}, MasterW,
							)
						}
					},
				})
			}
			dl.SetData(dl_rows)
		}
	}()

	// 底部声明
	github, _ := url.Parse("https://github.com/rock-rabbit/rain-service-gui")
	blog, _ := url.Parse("https://www.68wu.cn/")
	footer := container.NewCenter(container.NewHBox(widget.NewHyperlink("github", github), widget.NewHyperlink("blog", blog)))
	vbox.Add(footer)
	return vbox
}

// makeTools 创建工具栏
func makeTools() fyne.CanvasObject {
	btns := container.NewHBox()
	uriInput := module.NewUriEntry()

	btns.Add(uriInput)
	btns.Add(widget.NewButtonWithIcon("添加", theme.ContentAddIcon(), func() {
		uri := uriInput.Text
		matched, _ := regexp.MatchString(`^https?://.*?`, uri)
		if !matched {
			ErrorDialog(errors.New("url 解析错误。"))
			return
		}
		_, err := Rain.AddUri(&Uri{
			Uri:    uriInput.Text,
			Outdir: config.Outdir,
		})
		if err != nil {
			ErrorDialog(err)
		}
	}))
	return btns
}
