package service

import (
	"fyne.io/fyne/v2"
	"log"
	"memo/constant"
)

type Eat struct {
	pref fyne.Preferences
	data []string
}

func (e Eat) Save(restaurant string) {
	e.data = append(e.data, restaurant)
	e.pref.SetStringList(constant.Restaurant, e.data)
	log.Printf("Save restaurant:%s \n", restaurant)
}

func (e Eat) Restaurants() []string {
	return e.pref.StringList(constant.Restaurant)
}

func NewEat() *Eat {
	pref := fyne.CurrentApp().Preferences()
	data := pref.StringList(constant.Restaurant)

	return &Eat{
		pref: pref,
		data: data,
	}
}
