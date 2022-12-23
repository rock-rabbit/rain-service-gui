package module

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type UriEntry struct {
	widget.Entry
}

func NewUriEntry() *UriEntry {
	ue := &UriEntry{}
	ue.ExtendBaseWidget(ue)
	ue.SetPlaceHolder("请输入资源链接...")
	return ue
}

func (ue *UriEntry) MinSize() fyne.Size {
	return fyne.Size{Width: 800, Height: 0}
}
