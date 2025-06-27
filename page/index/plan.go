package index

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"memo/constant"
	"memo/page/index/plan"
)

type Plan struct {
	groups  []fyne.CanvasObject
	content fyne.CanvasObject
}

func NewPlan() *Plan {
	return &Plan{
		groups: []fyne.CanvasObject{
			plan.NewGroup(constant.Day).Content(),
			plan.NewGroup(constant.Week).Content(),
			plan.NewGroup(constant.Month).Content(),
		},
	}
}

func (p *Plan) Content() fyne.CanvasObject {
	if p.content != nil {
		return p.content
	}

	p.content = container.NewVBox(p.groups...)
	return p.content
}
