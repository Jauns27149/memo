package group

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/service"
)

type Add struct {
	button *widget.Button
	input  *widget.Entry
}

func (a *Add) Content() fyne.CanvasObject {
	a.button.OnTapped = func() {
		if a.input.Hidden {
			a.input.Show()
		} else {
			a.input.Hide()
		}
	}
	a.input.Hide()
	return container.NewVBox(a.button, a.input)
}

func NewAdd(group string) *Add {
	button := widget.NewButton("+", nil)
	input := widget.NewEntry()
	input.PlaceHolder = constant.MemoHolder
	input.OnSubmitted = func(text string) {
		service.PlanService.Save(group, text)
		input.Hide()
		input.SetText("")
	}

	return &Add{
		button: button,
		input:  input,
	}
}
