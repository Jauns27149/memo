package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(service.EatService.Data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox()
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			c := object.(*fyne.Container)
			c.RemoveAll()
			button := widget.NewButton("删除", func() {
				service.EatService.Delete(id)
				list.Refresh()
			})
			button.Importance = widget.HighImportance
			label := widget.NewLabel(service.EatService.Data[id])
			c.Add(label)
			c.Add(layout.NewSpacer())
			c.Add(button)
			list.SetItemHeight(id, c.MinSize().Height)
		},
	)

	return &Show{
		Head: widget.NewButton("X", nil),
		list: list,
	}
}
