package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"memo/model"
	"memo/service"
)

type Group struct {
	content fyne.CanvasObject
	span    string
	delete  bool
}

func (g *Group) Content() fyne.CanvasObject {
	if g.content != nil {
		return g.content
	}

	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(g.items())
		},
		func() fyne.CanvasObject {
			button := widget.NewButton("delete", nil)
			button.Importance = widget.DangerImportance
			return container.NewBorder(nil, nil, nil, button, nil)
		},
		func(i widget.ListItemID, object fyne.CanvasObject) {
			c := object.(*fyne.Container)
			c.Objects = c.Objects[:1]
			if g.delete {
				c.Objects[0].Show()
			} else {
				c.Objects[0].Hide()
			}
			button := widget.NewButton("", nil)
			button.Alignment = widget.ButtonAlignLeading
			items := g.items()
			button.SetText(items[i].Content)
			button.OnTapped = func() {
				service.PlanService.Finished(g.span, i)
				log.Println("finished item")
			}
			c.Add(button)
			if items[i].CreateTime.IsZero() {
				button.Enable()
			} else {
				button.Disable()
			}
		},
	)

	go g.listener(list)

	g.content = list
	return g.content
}

func (g *Group) listener(list *widget.List) {
	for {
		select {
		case <-service.PlanService.Update[g.span]:
			fyne.Do(func() {
				list.Refresh()
			})
		case value := <-service.PlanService.Delete[g.span]:
			fyne.Do(func() {
				g.delete = value
				list.Refresh()
			})
		}
	}
}

func (g *Group) items() []model.Item {
	for _, v := range service.PlanService.Items {
		if g.span == v.Span {
			return v.Items
		}
	}
	return []model.Item{}
}

func NewGroup(span string) *Group {
	return &Group{
		span: span,
	}
}
