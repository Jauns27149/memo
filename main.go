package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"memo/constant"
	"memo/page"
	"memo/service"
)

func main() {
	log.Println("memo is running")

	a := app.NewWithID(constant.AppName)
	w := a.NewWindow(constant.AppName)
	w.Resize(fyne.NewSize(400, 500))

	service.Boot()
	w.SetContent(page.NewIndex().Content())

	w.ShowAndRun()
}
