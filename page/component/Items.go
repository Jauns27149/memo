package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"memo/service"
)

type Items struct {
	content fyne.CanvasObject
}

func (i *Items) Content() fyne.CanvasObject {
	if i.content != nil {
		return i.content
	}

	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(service.MemoService.GetItems())
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", nil)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			button := o.(*widget.Button)
			button.Alignment = widget.ButtonAlignLeading
			button.SetText(service.MemoService.GetItems()[i].Content)
			if service.MemoService.GetItems()[i].Finished {
				button.Disable()
			} else {
				button.Enable()
			}

			button.OnTapped = func() {
				service.MemoService.Finished(i)
				list.Refresh()
			}
		},
	)

	i.content = list
	go listener(list)
	return i.content
}

func NewItems() *Items {
	return &Items{}
}

func listener(list *widget.List) {
	for {
		select {
		case <-service.MemoService.AddChan:
			fyne.Do(func() {
				list.Refresh()
			})
		}
	}
}
