package service

import (
	"fyne.io/fyne/v2"
	"log"
	"memo/constant"
	"memo/util"
	"strings"
	"time"
)

type Eat struct {
	pref fyne.Preferences
	Data []string
}

func (e *Eat) Save(restaurant string) {
	restaurant = strings.TrimSpace(restaurant)
	if restaurant == "" {
		log.Print("输入餐馆为空")
		return
	}

	for _, v := range e.Data {
		if v == restaurant {
			log.Printf("%v already exists\n", restaurant)
			return
		}
	}

	e.setRestaurants(append(e.Data, restaurant))
	log.Println("save restaurants finish")
}

func (e *Eat) LoadRestaurants() {
	e.Data = e.pref.StringList(constant.Restaurant)
	log.Printf("load restaurant success %v \n", e.Data)
}

func (e *Eat) Delete(id int) {
	list := e.Data
	e.setRestaurants(append(list[:id], list[id+1:]...))
	log.Printf("delete %v success\n", list[id])
}

func (e *Eat) setRestaurants(list []string) {
	e.pref.SetStringList(constant.Restaurant, list)
	log.Println("set restaurants finish")
}

func (e *Eat) RandRestaurant() string {
	data := e.Data

	for i, v := range data {
		if util.RestaurantHadTime(v) {
			t, _ := util.GetTime(v)
			if t.Add(3 * time.Hour).After(time.Now()) {
				return util.ClearTime(v) + " 已经选过了"
			} else {
				data[i] = util.ClearTime(v)
				e.setRestaurants(data)
			}
			break
		}
	}

	num := util.RandRestaurant(len(data))
	result := data[num]
	data[num] = data[num] + " " + time.Now().Format(time.DateTime)
	e.setRestaurants(data)
	return result
}

func NewEat() *Eat {
	pref := fyne.CurrentApp().Preferences()
	data := pref.StringList(constant.Restaurant)

	return &Eat{
		pref: pref,
		Data: data,
	}
}
