package eat

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/service"
)

type Add struct {
	button *widget.Button
	entry  *widget.Entry
}

func (a *Add) Content() fyne.CanvasObject {
	a.entry.Hide()
	a.button.OnTapped = func() {
		if a.entry.Hidden {
			a.entry.Show()
		} else {
			a.entry.Hide()
		}
	}

	head := container.NewHBox(layout.NewSpacer(), a.button)
	return container.NewVBox(head, a.entry)
}

func NewAdd() *Add {
	button := widget.NewButton(constant.AddChar, nil)
	entry := widget.NewEntry()
	entry.OnSubmitted = func(s string) {
		service.EatService.Save(s)
		entry.Hide()
	}

	return &Add{
		button: button,
		entry:  entry,
	}
}
