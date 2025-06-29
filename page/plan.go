package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"memo/constant"
	"memo/page/component"
)

type Plan struct {
	groups  []fyne.CanvasObject
	content fyne.CanvasObject
}

func NewPlan() *Plan {
	return &Plan{
		groups: []fyne.CanvasObject{
			component.NewGroup(constant.Day).Content(),
			component.NewGroup(constant.Week).Content(),
			component.NewGroup(constant.Month).Content(),
		},
	}
}

func (p *Plan) Content() fyne.CanvasObject {
	if p.content != nil {
		return p.content
	}

	p.content = container.NewGridWithRows(len(p.groups), p.groups...)
	return p.content
}
