package module

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type DownloadList struct {
	widget.List
}

type Row struct {
	UUID    string
	Status  string
	Title   string
	Progess float64
}

func NewDownloadList() *DownloadList {
	var data = []*Row{
		{UUID: "1", Status: "running", Title: "SimplifiedChinese1.jpg", Progess: 30},
		{UUID: "2", Status: "pause", Title: "SimplifiedChinese2.jpg", Progess: 40},
		{UUID: "3", Status: "finish", Title: "SimplifiedChinese3.jpg", Progess: 100},
	}
	dl := &DownloadList{
		List: widget.List{
			BaseWidget: widget.BaseWidget{},
			Length: func() int {
				return len(data)
			},
			CreateItem: func() fyne.CanvasObject {
				hbox := container.NewHBox()
				progress := NewMyProgress()
				progress.SetValue(50)
				ctl := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {})
				status := NewDownloadTitle("下载中", 50)

				hbox.Add(NewDownloadTitle("title....", 300))
				hbox.Add(progress)
				hbox.Add(status)
				hbox.Add(layout.NewSpacer())
				hbox.Add(ctl)

				ctl.SetIcon(theme.DeleteIcon())
				ctl.SetIcon(theme.MediaPlayIcon())
				ctl.SetIcon(theme.MediaPauseIcon())
				ctl.SetIcon(theme.ViewRefreshIcon())

				return hbox
			},
			UpdateItem: func(i widget.ListItemID, o fyne.CanvasObject) {
				objs := o.(*fyne.Container).Objects
				objs[0].(*DownloadTitle).SetText(data[i].Title)
				objs[1].(*MyProgress).SetValue(data[i].Progess)
				objs[2].(*DownloadTitle).SetText(StatusToChinese(data[i].Status))
				ctlIcon := theme.DeleteIcon()
				ctlText := "删除"
				switch data[i].Status {
				case "running":
					ctlIcon = theme.MediaPauseIcon()
					ctlText = "暂停"
				case "pause":
					ctlIcon = theme.MediaPlayIcon()
					ctlText = "继续"
				case "finish":
					ctlIcon = theme.ViewRefreshIcon()
					ctlText = "重试"
				case "error":
					ctlIcon = theme.ViewRefreshIcon()
					ctlText = "重试"
				}
				objs[4].(*widget.Button).SetIcon(ctlIcon)
				objs[4].(*widget.Button).SetText(ctlText)
			},
		},
	}
	dl.ExtendBaseWidget(dl)
	return dl
}

func (*DownloadList) MinSize() fyne.Size {
	return fyne.Size{Height: 500}
}

// StatusToChinese
func StatusToChinese(s string) string {
	switch s {
	case "running":
		return "下载中"
	case "pause":
		return "暂停"
	case "finish":
		return "完成"
	case "error":
		return "错误"
	}
	return "--"
}
