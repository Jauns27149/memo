package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/service"
)

type MemoAdd struct {
	Content fyne.CanvasObject
}

func NewMemoAdd() *MemoAdd {
	add := widget.NewButton("+", nil)
	item := widget.NewEntry()
	item.SetPlaceHolder(constant.MemoHolder)
	item.Hide()

	add.OnTapped = func() {
		if item.Visible() {
			item.Hide()
		} else {
			item.Show()
		}
	}
	item.OnSubmitted = func(s string) {
		item.Hide()
		service.MemoService.Add(s)
		item.SetText("")
	}

	return &MemoAdd{
		Content: container.NewVBox(add, item),
	}
}
