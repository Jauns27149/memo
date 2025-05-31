package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"memo/constant"
	"memo/layout"
	"memo/service"
)

func main() {
	log.Println("memo is running")

	w := app.NewWithID(constant.AppName).NewWindow(constant.AppName)
	w.Resize(fyne.NewSize(400, 600))
	service.MemoRun()
	service.PlanRun()

	var left, right *widget.Button
	var middle, content *fyne.Container
	left = widget.NewButton("便签", func() {
		left.Disable()
		right.Enable()
		content.Remove(service.Plan)
		content.Add(service.Memo)
	})
	left.Disable()

	right = widget.NewButton("计划表", func() {
		right.Disable()
		left.Enable()
		content.Remove(service.Memo)
		content.Add(service.Plan)
	})
	bottom := container.New(&layout.Menu{}, left, right)
	middle = service.Memo
	content = container.NewBorder(nil, bottom, nil, nil, middle)

	w.SetContent(content)
	w.ShowAndRun()
}
