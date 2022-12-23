package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rock-rabbit/rain-service-gui/module"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})

	w := a.NewWindow("Rain")
	w.SetContent(makeUI())
	w.SetFixedSize(true)
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
	vbox.Add(widget.NewLabel("下载列表"))
	vbox.Add(module.NewDownloadList())

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
	btns.Add(module.NewUriEntry())
	btns.Add(widget.NewButtonWithIcon("添加", theme.ContentAddIcon(), func() {}))
	return btns
}
