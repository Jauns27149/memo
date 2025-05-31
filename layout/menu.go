package layout

import (
	"fyne.io/fyne/v2"
)

type Menu struct{}

func (m *Menu) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 0 {
		return
	}

	unit := size.Width / float32(len(objects))

	for i, o := range objects {
		o.Resize(fyne.NewSize(unit, size.Height))
		o.Move(fyne.NewPos(float32(i)*unit, 0))
	}
}

func (m *Menu) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var minWidth, minHeight float32
	for _, o := range objects {
		minWidth += o.Size().Width
		minHeight = max(minHeight, o.MinSize().Height)
	}
	return fyne.NewSize(minWidth, minHeight)
}
