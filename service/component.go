package service

import "fyne.io/fyne/v2"

type Component interface {
	Content() fyne.CanvasObject
}
