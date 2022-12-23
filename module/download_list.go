package module

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type DownloadList struct {
	widget.List
	data []*Row
}

type Row struct {
	UUID    string
	Status  string
	Title   string
	Progess float64
	// TotalLength 文件总大小
	TotalLength int64
	// CompletedLength 已下载的文件大小
	CompletedLength int64
	// DownloadSpeed 每秒下载字节数
	DownloadSpeed int64
	Click         func(c string)
}

func NewDownloadList() *DownloadList {
	// var data = []*Row{
	// 	{UUID: "1", Status: "running", Title: "SimplifiedChinese1.jpg", Progess: 30},
	// 	{UUID: "2", Status: "pause", Title: "SimplifiedChinese2.jpg", Progess: 40},
	// 	{UUID: "3", Status: "finish", Title: "SimplifiedChinese3.jpg", Progess: 100},
	// }
	data := make([]*Row, 0)
	dl := &DownloadList{}
	dl.data = data
	dl.Length = func() int {
		return len(dl.GetData())
	}
	dl.CreateItem = func() fyne.CanvasObject {
		hbox := container.NewHBox()
		progress := NewMyProgress()
		progress.SetValue(50)
		progress.TextFormatter = func() string {
			return "100%"
		}
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
	}
	dl.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
		data := dl.GetData()[i]

		objs := o.(*fyne.Container).Objects
		objs[0].(*DownloadTitle).SetText(data.Title)
		objs[1].(*MyProgress).SetValue(data.Progess)
		objs[1].(*MyProgress).TextFormatter = func() string {
			if data.Status == "finish" || data.Status == "error" {
				return fmt.Sprintf(
					"%d%%",
					int(data.Progess),
				)
			}
			return fmt.Sprintf(
				"%d%%\t\t%s/s",
				int(data.Progess),
				FormatFileSize(data.DownloadSpeed),
			)
		}
		objs[2].(*DownloadTitle).SetText(StatusToChinese(data.Status))
		ctlIcon := theme.DeleteIcon()
		ctlText := "删除"
		switch data.Status {
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
		objs[4].(*widget.Button).OnTapped = func() {
			data.Click(ctlText)
		}
	}
	dl.ExtendBaseWidget(dl)
	return dl
}
func (dl *DownloadList) GetData() []*Row {
	return dl.data
}

func (dl *DownloadList) SetData(data []*Row) {
	dl.data = data
	dl.Refresh()
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

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) string {
	var sizef = float64(fileSize)
	if sizef <= 0 {
		return "0.00 B"
	}
	if sizef < 1024 {
		return fmt.Sprintf("%.2f B", sizef/float64(1))
	} else if sizef < 1048576 {
		return fmt.Sprintf("%.2f Kib", sizef/float64(1024))
	} else if sizef < 1073741824 {
		return fmt.Sprintf("%.2f Mib", sizef/float64(1048576))
	} else if sizef < 1099511627776 {
		return fmt.Sprintf("%.2f Gib", sizef/float64(1073741824))
	} else if sizef < 1125899906842624 {
		return fmt.Sprintf("%.2f Tib", sizef/float64(1099511627776))
	} else {
		return fmt.Sprintf("%.2f Eib", sizef/float64(1125899906842624))
	}
}
