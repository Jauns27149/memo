package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"memo/constant"
	"memo/page/component"
)

type Index struct {
	memo fyne.CanvasObject
	plan fyne.CanvasObject
	eat  fyne.CanvasObject

	content *fyne.Container
	bottom  *component.Bottom
	current int
}

func NewIndex() *Index {
	return &Index{
		memo:   NewMemo().Content(),
		plan:   NewPlan().Content(),
		eat:    NewEat().Content(),
		bottom: component.NewBottom(),
	}
}

func (i *Index) Content() fyne.CanvasObject {
	texts := i.bottom.Texts
	contentMap := make(map[string]fyne.CanvasObject, len(texts))
	contentMap[constant.Memo] = i.memo
	contentMap[constant.Plan] = i.plan
	contentMap[constant.Eat] = i.eat

	buttons := i.bottom.Buttons
	for ii, _ := range texts {
		oneself := ii
		buttons[ii].OnTapped = func() {
			buttons[oneself].Disable()
			buttons[i.current].Enable()
			i.content.Remove(contentMap[texts[i.current]])
			i.content.Add(contentMap[texts[oneself]])
			i.current = oneself
		}
	}
	buttons[i.current].Disable()

	i.content = container.NewBorder(nil, i.bottom.Content(), nil, nil, i.memo)
	return i.content
}
