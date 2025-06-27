package service

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"log"
	"memo/constant"
	"memo/util"
	"time"
)

type Plan struct {
	pref fyne.Preferences
	Data map[string]binding.StringList
}

func (p Plan) LoadData(group string) []string {
	return []string{}
}

func (p Plan) Save(group string, text string) {
	data := p.Data[group]
	err := data.Append(text)
	if err != nil {
		log.Panicln(err)
	}
	items, err := data.Get()
	if err != nil {
		log.Panicln(err)
	}
	p.pref.SetStringList(group, items)
	log.Printf("save items successfully, key: %v, value: %v\n", group, items)
}

func (p Plan) Done(ii int, span string, text string) {
	text = fmt.Sprintf("%s %s", text, time.Now().Format(time.DateTime))
	items, err := p.Data[span].Get()
	if err != nil {
		log.Panicln(err)
	}
	items[ii] = text
	err = p.Data[span].Remove(items[0])
	if err != nil {
		log.Panicln(err)
	}

	util.SortPlanItems(items)
	err = p.Data[span].Set(items)
	if err != nil {
		log.Panicln(err)
	}
}

func NewPlan() *Plan {
	pref := fyne.CurrentApp().Preferences()
	dataMap := make(map[string]binding.StringList)
	for _, v := range []string{constant.Day, constant.Week, constant.Month} {
		dataMap[v] = binding.BindPreferenceStringList(v, pref)
	}
	return &Plan{
		pref: pref,
		Data: dataMap,
	}
}
