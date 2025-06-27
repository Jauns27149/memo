package index

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"memo/page/index/eat"
)

type Eat struct {
	add    fyne.CanvasObject
	random fyne.CanvasObject
}

func (e *Eat) Content() fyne.CanvasObject {
	return container.NewBorder(e.add, nil, nil, nil, e.random)
}

func NewEat() *Eat {
	return &Eat{
		add:    eat.NewAdd().Content(),
		random: eat.NewRandom().Content(),
	}
}
