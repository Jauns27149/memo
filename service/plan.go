package service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/layout"
	"strings"
	"time"
)

var Plan *fyne.Container

func PlanRun() {
	item := make([]fyne.CanvasObject, 3)
	for i, v := range []string{"日", "周", "月"} {
		item[i] = createUnit(v, i)
	}
	Plan = container.NewVBox(item...)
}

// TODO 计划界面待完善
func createUnit(value string, i int) *fyne.Container {
	entry := widget.NewEntry()
	sure := widget.NewButton("确定", nil)
	cancel := widget.NewButton("取消", nil)
	c := container.New(&layout.Item{}, entry, sure, cancel)
	c.Hide()
	cancel.OnTapped = func() {
		c.Hide()
	}

	preferences := fyne.CurrentApp().Preferences()
	sure.OnTapped = func() {
		list := preferences.StringList(value)
		preferences.SetStringList(value, append(list, entry.Text))

		objects := Plan.Objects[i].(*fyne.Container).Objects
		button := widget.NewButton(entry.Text, nil)
		button.OnTapped = func() {
			button.Disable()
			t := preferences.StringList(value)
			t[len(list)] = t[len(list)] + "-" + time.Now().Format("2006-01-02")
			preferences.SetStringList(value, t)
		}

		Plan.Objects[i].(*fyne.Container).Objects = append(objects, button)
		entry.SetText("")
		c.Hide()
	}

	button := widget.NewButton(constant.AddChar, func() {
		c.Show()
	})

	label := widget.NewLabel(value)
	day := container.New(&layout.Item{}, label, button)
	dayList := preferences.StringList(value)
	// 格式：thing-done_time
	temp := make([]fyne.CanvasObject, len(dayList)+2)
	temp[0] = day
	temp[1] = c
	for i, v := range dayList {
		values := strings.Split(v, "-")
		b := widget.NewButton(values[0], nil)
		b.OnTapped = func() {
			b.Disable()
			t := preferences.StringList(value)
			t[i] = t[i] + "-" + time.Now().Format("2006-01-02")
			preferences.SetStringList(value, t)
		}
		if len(values) > 1 {
			b.Disable()
		}
		temp[i+2] = b
	}
	v := container.NewVBox(temp...)
	v.Resize(v.MinSize())
	return v
}
