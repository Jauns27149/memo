package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"memo/page/component"
	"memo/service"
)

type Eat struct {
	contents []fyne.CanvasObject
	menu     []*widget.Button
	current  int
}

func (e *Eat) Content() fyne.CanvasObject {
	var content *fyne.Container
	objects := make([]fyne.CanvasObject, len(e.contents)+2)
	objects[0], objects[len(objects)-1] = layout.NewSpacer(), layout.NewSpacer()
	box := container.NewVBox(objects...)
	for i, button := range e.menu {
		objects[i+1] = button
		ii := i
		button.OnTapped = func() {
			content.RemoveAll()
			switch ii {
			case 0:
				b := e.contents[ii].(*widget.Button)
				b.SetText(service.EatService.RandRestaurant())
				b.OnTapped = func() {
					content.Objects = append(content.Objects, box)
				}
				content.Add(container.NewCenter(e.contents[ii]))

			case 1:
				b := e.contents[ii].(*fyne.Container).Objects[1].(*widget.Button)
				b.Alignment = widget.ButtonAlignTrailing
				b.OnTapped = func() {
					content.Objects = []fyne.CanvasObject{box}
				}
				content.Add(e.contents[ii])

			case 2:
				e.contents[ii].(*widget.Entry).OnSubmitted = func(restaurant string) {
					service.EatService.Save(restaurant)
					content.Objects = append(content.Objects, box)
				}
				content.Add(container.NewVBox(layout.NewSpacer(), e.contents[ii], layout.NewSpacer()))

			}
			e.current = ii
		}
	}

	content = container.NewBorder(nil, nil, nil, nil, box)
	return content
}

func NewEat() *Eat {
	texts := []string{"选择", "查看", "新增"}
	buttons := make([]*widget.Button, len(texts))
	for i, text := range texts {
		buttons[i] = widget.NewButton(text, nil)
	}
	content := []fyne.CanvasObject{
		widget.NewButton("", nil),
		component.NewShow().Content(),
		component.NewAdd().Content(),
	}

	return &Eat{
		menu:     buttons,
		contents: content,
	}
}
