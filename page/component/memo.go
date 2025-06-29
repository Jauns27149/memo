package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"memo/page/index/memo"
)

type Memo struct {
	content fyne.CanvasObject
	items   fyne.CanvasObject
	add     fyne.CanvasObject
}

func NewMemo() *Memo {
	return &Memo{
		add:   memo.NewAdd().Content,
		items: memo.NewItems().Content(),
	}
}

func (m *Memo) Content() fyne.CanvasObject {
	if m.content != nil {
		return m.content
	}

	m.content = container.NewVBox(m.add, m.items)
	return m.content
}
