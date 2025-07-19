package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/page/component"
	"memo/service"
)

type Plan struct {
	groups  []fyne.CanvasObject
	content fyne.CanvasObject
	head    fyne.CanvasObject
	current int
}

func NewPlan() *Plan {
	items := service.PlanService.Items
	groups := make([]fyne.CanvasObject, len(items))
	for i, v := range items {
		groups[i] = component.NewGroup(v.Span).Content()
	}

	list := make([]string, len(items))
	for i, v := range items {
		list[i] = v.Span
	}
	plan := &Plan{
		groups: groups,
	}

	selectWidget := widget.NewSelect(list, func(s string) {
		for i, v := range items {
			if v.Span == s && plan.content != nil {
				c := plan.content.(*fyne.Container)
				c.Remove(groups[plan.current])
				c.Add(groups[i])
				plan.current = i
				break
			}
		}
	})

	selectWidget.SetSelected(constant.Day)
	entry := widget.NewEntry()
	entry.OnSubmitted = func(s string) {
		service.PlanService.Add(selectWidget.Selected, s)
	}
	entry.Hide()
	add := widget.NewButton(constant.AddChar, func() {
		if entry.Hidden {
			entry.Show()
		} else {
			entry.Hide()
		}
	})
	add.Importance = widget.WarningImportance
	flag := false
	manage := widget.NewButton("管理", func() {
		if flag {
			flag = false
		} else {
			flag = true
		}
		service.PlanService.Delete[selectWidget.Selected] <- flag
	})

	plan.head = container.NewVBox(container.NewHBox(selectWidget, layout.NewSpacer(), manage, add), entry)

	return plan
}

func (p *Plan) Content() fyne.CanvasObject {
	if p.content != nil {
		return p.content
	}

	p.content = container.NewBorder(p.head, nil, nil, nil)
	if p.groups != nil && len(p.groups) > 0 {
		p.content.(*fyne.Container).Add(p.groups[0])
	}

	return p.content
}
