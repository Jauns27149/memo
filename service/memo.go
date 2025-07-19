package service

import (
	"fyne.io/fyne/v2"
	"log"
	"memo/constant"
	"memo/convert"
	"memo/model"
	"memo/util"
	"strings"
	"time"
)

type Memo struct {
	preferences fyne.Preferences
	items       []model.MemoItem
	AddChan     chan bool
}

func (d *Memo) GetItems() []model.MemoItem {
	util.SortMemo(d.items)
	return d.items
}

func newData() *Memo {
	preferences := fyne.CurrentApp().Preferences()
	memo := preferences.StringList(constant.Memo)
	item, err := convert.RowsToMemo(memo)
	if err != nil {
		log.Fatal(err)
	}
	return &Memo{
		preferences: preferences,
		items:       item,
		AddChan:     make(chan bool, 1),
	}
}

func (d *Memo) Add(item string) {
	memoItem := model.MemoItem{Item: model.Item{Content: strings.TrimSpace(item), CreateTime: time.Now()}, Finished: false}
	d.items = append(d.items, memoItem)
	util.SortMemo(d.items)
	d.preferences.SetStringList(constant.Memo, convert.MemosToRows(d.items))
	d.AddChan <- true
}

func (d *Memo) Finished(index int) {
	d.items[index].Finished = true
	d.preferences.SetStringList(constant.Memo, convert.MemosToRows(d.items))
}

func (d *Memo) Delete(index int) {
	d.items = append(d.items[:index], d.items[index+1:]...)
	d.preferences.SetStringList(constant.Memo, convert.MemosToRows(d.items))
}
