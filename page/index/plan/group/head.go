package group

import (
	"fyne.io/fyne/v2/widget"
)

type Head struct {
	Content *widget.Button
}

func NewHead(span string) *Head {
	button := widget.NewButton(span, nil)
	button.Alignment = widget.ButtonAlignLeading
	return &Head{
		Content: button,
	}
}
