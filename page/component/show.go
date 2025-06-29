package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/service"
)

type Show struct {
	Head *widget.Button
	list *widget.List
}

func (s *Show) Content() fyne.CanvasObject {
	return container.NewBorder(s.Head, nil, nil, nil, s.list)
}

func NewShow() *Show {
	list := widget.NewList(
		func() int {
			return len(service.EatService.Data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(service.EatService.Data[id])
		},
	)

	return &Show{
		Head: widget.NewButton("X", nil),
		list: list,
	}
}
