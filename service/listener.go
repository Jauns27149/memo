package service

import (
	"fyne.io/fyne/v2"
	"log"
)

func listenerPref() {
	fyne.CurrentApp().Preferences().AddChangeListener(func() {
		EatService.LoadRestaurants()
		//PlanService.LoadData()

		if w := fyne.CurrentApp().Driver().AllWindows(); len(w) > 0 {
			w[0].Content().Refresh()
		}
		log.Printf("listenter data change %v \n", EatService.Data)
	})
}
