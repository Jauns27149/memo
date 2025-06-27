package eat

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"memo/constant"
	"memo/service"
	"memo/util"
)

type Random struct {
	choice *widget.Button
	show   *canvas.Text
}

func (r *Random) Content() fyne.CanvasObject {
	r.show.Hide()
	r.choice.OnTapped = func() {
		restaurants := service.EatService.Restaurants()
		restaurant := util.RandRestaurant(restaurants)
		r.show.Text = restaurant
		r.show.Show()
	}

	return container.NewVBox(layout.NewSpacer(), r.choice, r.show, layout.NewSpacer())
}

func NewRandom() *Random {
	choice := widget.NewButton(constant.Choice, nil)
	show := canvas.NewText("", nil)
	show.Alignment = fyne.TextAlignCenter

	return &Random{
		choice: choice,
		show:   show,
	}
}
