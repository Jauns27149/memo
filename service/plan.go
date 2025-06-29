package service

import (
	"fyne.io/fyne/v2"
	"log"
	"memo/constant"
	"memo/util"
	"time"
)

type Plan struct {
	pref fyne.Preferences

	Day   []string
	Week  []string
	Month []string
}

func (p *Plan) LoadData() {
	p.Day = p.loadItems(constant.Day)
	p.Week = p.loadItems(constant.Week)
	p.Month = p.loadItems(constant.Month)
	log.Println("load plan data finished")
}

func (p *Plan) Save(span string, text string) {
	list := p.pref.StringList(span)
	p.set(span, append(list, text))
	log.Printf("save items successfully")
}

func (p *Plan) Done(ii int, span string) {
	list := p.Item(span)
	list[ii] = list[ii] + " " + time.Now().Format(time.DateTime)
	p.set(span, list)
	log.Printf("item done successfully")
}

func (p *Plan) set(span string, list []string) {
	// 确保监听函数触发，数量没有变化的时候不会触发
	p.pref.SetStringList(span, []string{})
	p.pref.SetStringList(span, list)
}

func (p *Plan) loadItems(span string) []string {
	list := p.pref.StringList(span)
	util.SortPlanItems(list)
	return list
}

func (p *Plan) Item(span string) []string {
	var list []string
	switch span {
	case constant.Day:
		list = p.Day
	case constant.Week:
		list = p.Week
	case constant.Month:
		list = p.Month
	}
	return list
}

func NewPlan() *Plan {
	pref := fyne.CurrentApp().Preferences()
	plan := &Plan{
		pref: pref,
	}
	plan.LoadData()
	return plan
}
