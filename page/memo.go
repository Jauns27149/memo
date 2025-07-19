package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"memo/page/component"
)

type Memo struct {
	content fyne.CanvasObject
	items   fyne.CanvasObject
	add     fyne.CanvasObject
}

func NewMemo() *Memo {
	return &Memo{
		add:   component.NewMemoAdd().Content,
		items: component.NewItems().Content(),
	}
}

func (m *Memo) Content() fyne.CanvasObject {
	if m.content != nil {
		return m.content
	}

	m.content = container.NewBorder(m.add, nil, nil, nil, m.items)
	return m.content
}
