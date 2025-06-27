package index

import "C"
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
)

type Bottom struct {
	Texts   []string
	Buttons []*widget.Button
}

func (b *Bottom) Content() *fyne.Container {
	objects := make([]fyne.CanvasObject, len(b.Buttons))
	for i, button := range b.Buttons {
		objects[i] = button
	}
	return container.NewGridWithColumns(len(b.Texts), objects...)
}

func NewBottom() *Bottom {
	texts := []string{constant.Memo, constant.Plan, constant.Eat}
	buttons := make([]*widget.Button, len(texts))
	for i, text := range texts {
		buttons[i] = widget.NewButton(text, nil)
	}
	return &Bottom{
		Texts:   texts,
		Buttons: buttons,
	}
}
