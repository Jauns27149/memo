package util

import (
	"memo/model"
	"sort"
)

func SortMemo(memo []model.MemoItem) {
	sort.Slice(memo, func(i, j int) bool {
		if memo[j].Finished != memo[i].Finished {
			return memo[j].Finished
		}
		return memo[i].CreateTime.After(memo[j].CreateTime)
	})
}

func SortPlanItems(items []model.Item) {
	end := len(items)
	for i := 0; i < end; i++ {
		if !items[i].CreateTime.IsZero() && i != end-1 {
			items[i], items[end-1] = items[end-1], items[i]
			i--
			end--
		}
	}
}
