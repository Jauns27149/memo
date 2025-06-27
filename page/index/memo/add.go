package memo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/service"
)

type Add struct {
	Content fyne.CanvasObject
}

func NewAdd() *Add {
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

	return &Add{
		Content: container.NewVBox(add, item),
	}
}
