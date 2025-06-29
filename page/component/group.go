package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"memo/constant"
	"memo/service"
	"memo/util"
)

type Span = string

type Group struct {
	head  []fyne.CanvasObject
	items *widget.List
	add   *widget.Entry
	span  Span
}

func (g *Group) Content() fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, len(g.head))
	for i, item := range g.head {
		objects[i] = item
	}
	border := container.NewBorder(nil, nil, nil, objects[1], objects[0])

	g.add.Hide()
	g.add.OnSubmitted = func(s string) {
		service.PlanService.Save(g.span, s)
		g.add.Hide()
		g.items.Refresh()
	}

	head := container.NewVBox(border, g.add)
	return container.NewBorder(head, nil, nil, nil, g.items)
}

func NewGroup(span Span) *Group {
	add := widget.NewEntry()

	label := widget.NewLabel(span)
	button := widget.NewButton(constant.AddChar, nil)
	button.OnTapped = func() {
		add.Show()
	}
	head := []fyne.CanvasObject{label, button}
	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(service.PlanService.Item(span))
		},
		func() fyne.CanvasObject {
			b := widget.NewButton("", nil)
			b.Alignment = widget.ButtonAlignLeading
			return b
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			button := item.(*widget.Button)
			s := service.PlanService.Item(span)[id]
			button.SetText(s)
			button.OnTapped = func() {
				service.PlanService.Done(id, span)
				list.Refresh()
			}
			if util.HadTime(s) {
				time, err := util.GetTime(s)
				if err != nil {
					log.Panicln(err.Error())
				}
				updateTime := util.LatePan(span, time)
				if time.Before(updateTime) {
					service.PlanService.Save(span, util.ClearTime(s))
				} else {
					button.SetText(util.ClearTime(s))
					button.Disable()
				}
			}
		},
	)

	return &Group{
		head:  head,
		items: list,
		add:   add,
		span:  span,
	}
}
