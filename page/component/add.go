package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Add struct {
}

func (a *Add) Content() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("write a restaurant name")
	return entry
}

func NewAdd() *Add {
	return &Add{}
}
