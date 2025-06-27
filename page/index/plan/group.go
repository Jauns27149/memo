package plan

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/page/index/plan/group"
)

type Group struct {
	head   *widget.Button
	center fyne.CanvasObject
	add    fyne.CanvasObject
	span   string
}

func NewGroup(span string) *Group {
	return &Group{
		head:   group.NewHead(span).Content,
		add:    group.NewAdd(span).Content(),
		center: group.NewCenter(span).Content(),
		span:   span,
	}
}

func (g *Group) Content() fyne.CanvasObject {
	g.head.OnTapped = func() {
		if g.center.Visible() && g.add.Visible() {
			g.center.Hide()
			g.add.Hide()
		} else {
			g.center.Show()
			g.add.Show()
		}
	}

	if g.span != constant.Day {
		g.center.Hide()
		g.add.Hide()
	}
	return container.NewVBox(g.head, g.add, g.center)
}
