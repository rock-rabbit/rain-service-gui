package module

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type DownloadTitle struct {
	widget.Label
	width float32
}

func NewDownloadTitle(t string, size float32) *DownloadTitle {
	ue := &DownloadTitle{}
	ue.ExtendBaseWidget(ue)
	ue.SetText(t)
	ue.width = size
	return ue
}

func (ue *DownloadTitle) MinSize() fyne.Size {
	return fyne.Size{Width: ue.width, Height: 30}
}
