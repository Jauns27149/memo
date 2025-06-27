package service

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"log"
	"memo/constant"
	"memo/util"
	"strings"
	"time"
)

type Data struct {
	preferences fyne.Preferences
	Memo        binding.StringList
}

func newData() *Data {
	preferences := fyne.CurrentApp().Preferences()

	return &Data{
		preferences: preferences,
		Memo:        binding.BindPreferenceStringList(constant.Memo, preferences),
	}
}

func (d *Data) Add(item string) {
	item = fmt.Sprintf("%s %s %s", item, time.Now().Format(time.DateTime), constant.Doing)
	err := d.Memo.Append(item)
	if err != nil {
		log.Panicln(err)
	}

	util.SortMemo(d.Memo)
	d.save()
}

func (d *Data) DoneItem(index int) {
	value, err := d.Memo.GetValue(index)
	if err != nil {
		log.Panicln(err)
	}

	err = d.Memo.SetValue(index, strings.ReplaceAll(value, constant.Doing, constant.Done))
	if err != nil {
		log.Panicln(err)
	}
	util.SortMemo(d.Memo)
	d.save()
}

func (d *Data) save() {
	items, err := d.Memo.Get()
	if err != nil {
		log.Panicln(err)
	}
	d.preferences.SetStringList(constant.Memo, items)
}
