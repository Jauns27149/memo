package service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/layout"
)

var Memo *fyne.Container

func MemoRun() {
	var plus *fyne.Container
	entry := widget.NewEntry()
	sure := widget.NewButton("确定", func() {
		preferences := fyne.CurrentApp().Preferences()
		o := Memo.Objects
		button := widget.NewButton(entry.Text, nil)
		button.OnTapped = func() {
			Memo.Remove(button)
			Memo.Add(button)
			button.Disable()
			preferences.SetStringList(
				"done", append([]string{entry.Text}, preferences.StringList("done")...))
		}

		Memo.Objects = append([]fyne.CanvasObject{o[0], o[1], button}, o[2:]...)
		preferences.SetStringList(
			"todo", append([]string{entry.Text}, preferences.StringList("todo")...))

		entry.Text = ""
		Memo.Refresh()
		plus.Hide()
	})
	cancel := widget.NewButton("取消", func() {
		plus.Hide()
	})
	plus = container.New(&layout.Item{}, entry, sure, cancel)
	plus.Hide()

	canvas := []fyne.CanvasObject{widget.NewButton(constant.AddChar, func() {
		plus.Show()
	}), plus}

	preference := fyne.CurrentApp().Preferences()
	todos := preference.StringList("todo")
	dones := preference.StringList("done")
	for i, v := range todos {
		button := widget.NewButton(v, nil)
		button.OnTapped = func() {
			Memo.Remove(button)
			Memo.Add(button)
			button.Disable()
			preference.SetStringList("done", append(dones, v))
			preference.SetStringList("todo", append(todos[:i], todos[i+1:]...))
		}
		canvas = append(canvas, button)
	}
	for _, v := range dones {
		b := widget.NewButton(v, nil)
		canvas = append(canvas, b)
		b.Disable()
	}

	Memo = container.NewVBox(canvas...)
}
