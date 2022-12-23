package module

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MyProgress struct {
	widget.ProgressBar
}

func NewMyProgress() *MyProgress {
	mp := &MyProgress{}
	mp.ExtendBaseWidget(mp)
	mp.Min = 0
	mp.Max = 100
	return mp
}

func (mp *MyProgress) MinSize() fyne.Size {
	return fyne.Size{Width: 400, Height: 0}
}
