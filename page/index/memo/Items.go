package memo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
	"memo/constant"
	"memo/service"
	"regexp"
	"strings"
)

var items = &Items{content: container.NewVBox()}

type Items struct {
	content *fyne.Container
}

func NewItems() *Items {
	listener()
	return items
}

func (i *Items) Content() fyne.CanvasObject {
	i.content = container.NewVBox(itemsButton()...)
	return i.content
}

func itemsButton() []fyne.CanvasObject {
	memo, err := service.MemoService.Memo.Get()
	if err != nil {
		log.Panicln(err.Error())
	}

	objects := make([]fyne.CanvasObject, 0, len(memo))
	for i, item := range memo {
		reg := regexp.MustCompile("^[^ ]+ ")
		text := reg.FindString(item)
		button := widget.NewButton(text, nil)
		button.Alignment = widget.ButtonAlignLeading
		if strings.Contains(item, constant.Done) {
			button.Disable()
		}

		index := i
		button.OnTapped = func() {
			service.MemoService.DoneItem(index)
			button.Disable()
		}
		objects = append(objects, button)
	}
	return objects
}

func listener() {
	service.MemoService.Memo.AddListener(binding.NewDataListener(func() {
		memo := itemsButton()
		items.content.RemoveAll()
		for _, item := range memo {
			items.content.Add(item)
		}
		log.Println("update memo items success")
	}))
}
