package service

import (
	"fyne.io/fyne/v2"
	"log"
	"memo/constant"
	"memo/convert"
	"memo/model"
	"memo/util"
	"time"
)

type Plan struct {
	pref   fyne.Preferences
	Items  []model.PlanItem
	Update map[string]chan bool
	Delete map[string]chan bool
}

func (p *Plan) Add(span, content string) {
	for i, item := range p.Items {
		if item.Span == span {
			if content != "" {
				p.Items[i].Items = append(p.Items[i].Items, model.Item{Content: content})
			}
			p.pref.SetStringList(span, convert.PlanItemsToRows(p.Items[i].Items))
			p.Update[span] <- true
			break
		}
	}
}

func (p *Plan) Finished(span string, index int) {
	for i, item := range p.Items {
		if item.Span == span {
			p.Items[i].Items[index].CreateTime = time.Now()
			p.pref.SetStringList(span, convert.PlanItemsToRows(item.Items))
			log.Println("finished item")
			p.Update[span] <- true
			log.Println("update chan item")
			break
		}
	}
}

func NewPlan() *Plan {
	pref := fyne.CurrentApp().Preferences()
	span := []string{constant.Day, constant.Week, constant.Month}
	items := make([]model.PlanItem, len(span))
	update := make(map[string]chan bool, len(span))
	deleteChan := make(map[string]chan bool, len(span))
	for i, v := range span {
		update[v] = make(chan bool, 1)
		deleteChan[v] = make(chan bool, 1)
		list := pref.StringList(v)
		item, flag, err := convert.RowsToPlanItem(v, list)
		if err != nil {
			log.Fatal(err)
		}
		if flag {
			pref.SetStringList(v, convert.PlanItemsToRows(item.Items))
		}
		util.SortPlanItems(item.Items)
		items[i] = item
	}

	plan := &Plan{
		pref:   pref,
		Items:  items,
		Update: update,
		Delete: deleteChan,
	}
	return plan
}
