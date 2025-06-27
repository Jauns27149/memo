package util

import (
	"fyne.io/fyne/v2/data/binding"
	"log"
	"memo/constant"
	"sort"
	"strings"
)

func SortMemo(memo binding.StringList) {
	items, err := memo.Get()
	if err != nil {
		log.Panicln(err)
	}

	var done, doing []string
	for _, item := range items {
		if strings.Contains(item, constant.Done) {
			done = append(done, item)
		} else {
			doing = append(doing, item)
		}
	}

	sortMemo(done)
	sortMemo(doing)

	// 改变值不触发更新，这个只是触发更新，没有实际作用
	err = memo.Remove(items[0])
	if err != nil {
		log.Panicln(err)
	}

	err = memo.Set(append(doing, done...))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("sort items success")
}

func sortMemo(items []string) {
	sort.Slice(items, func(i, j int) bool {
		return findTime(items[i]).After(findTime(items[j]))
	})
}

func SortPlanItems(items []string) {
	end := len(items)
	for i := 0; i < end; i++ {
		if HadTime(items[i]) && i != end-1 {
			items[i], items[end-1] = items[end-1], items[i]
			i--
			end--
		}
	}
}
