package main

import (
	"errors"
	"net/url"
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
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})

	w := a.NewWindow("Rain")
	w.SetContent(makeUI())
	w.SetFixedSize(true)

	// 错误提示框
	ErrorDialog = func(err error) {
		dialog.ShowError(err, w)
	}

	w.ShowAndRun()
}

// makeUI 创建内容
func makeUI() fyne.CanvasObject {
	vbox := container.NewVBox()

	// 标题
	hbox := container.NewHBox(
		widget.NewLabel("Rain 下载器"),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("设置", theme.SettingsIcon(), func() {}),
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
						case "继续", "重试":
							Rain.Start(uuid)
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
		_, err := Rain.AddUri(&Uri{
			Uri: uriInput.Text,
		})
		if err != nil {
			ErrorDialog(errors.New("url 解析错误。"))
		}
	}))
	return btns
}
