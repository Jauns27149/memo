package group

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
	"memo/service"
	"memo/util"
)

type Center struct {
	group   string
	content *fyne.Container
	data    binding.StringList
}

func NewCenter(group string) *Center {
	data := service.PlanService.Data[group]
	return &Center{
		group: group,
		data:  data,
	}
}

func (g *Center) Content() fyne.CanvasObject {
	if g.content != nil {
		return g.content
	}

	g.content = container.NewVBox(g.objects()...)
	g.listener()
	return g.content
}

func (g *Center) objects() []fyne.CanvasObject {
	items, err := g.data.Get()
	if err != nil {
		log.Panicln(err)
	}
	util.SortPlanItems(items)
	objects := make([]fyne.CanvasObject, len(items))
	for i, item := range items {
		value := util.GetItemNoTime(item)
		button := widget.NewButton(value, nil)
		if util.HadTime(item) {
			button.Disable()
		}
		button.Alignment = widget.ButtonAlignLeading
		ii := i
		button.OnTapped = func() {
			button.Disable()
			service.PlanService.Done(ii, g.group, button.Text)
		}
		objects[i] = button
	}
	return objects
}

func (g *Center) listener() {
	g.data.AddListener(binding.NewDataListener(func() {
		g.content.RemoveAll()
		for _, object := range g.objects() {
			g.content.Add(object)
		}
		g.content.Refresh()
		log.Println("listener successful")
	}))
}
