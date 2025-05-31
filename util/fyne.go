package util

import "fyne.io/fyne/v2"

func GetTopWindow() fyne.Window {
	return fyne.CurrentApp().Driver().AllWindows()[0]
}

func GetContent() *fyne.Container {
	content := fyne.CurrentApp().Driver().AllWindows()[0].Content()
	return content.(*fyne.Container).Objects[0].(*fyne.Container)
}
