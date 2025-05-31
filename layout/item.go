package layout

import (
	"fyne.io/fyne/v2"
	"log"
)

type Item struct{}

func (item *Item) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 0 {
		return
	}

	currentWidth := size.Width
	for i := len(objects) - 1; i > 0; i-- {
		currentWidth -= objects[i].MinSize().Width
		objects[i].Resize(fyne.NewSize(objects[i].MinSize().Width, size.Height))
		objects[i].Move(fyne.NewPos(currentWidth, 0))
	}

	objects[0].Resize(fyne.NewSize(currentWidth, size.Height))
	objects[0].Move(fyne.NewPos(0, 0))
	log.Printf("size:%v,position:%v", objects[0].Size(), objects[0].Position())

}

func (item *Item) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var minWidth, minHeight float32
	for _, o := range objects {
		minWidth += o.Size().Width
		minHeight = max(minHeight, o.MinSize().Height)
	}
	return fyne.NewSize(minWidth, minHeight)
}
