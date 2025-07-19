package model

import "time"

type MemoItem struct {
	Item
	Finished bool
}

type Item struct {
	Content    string
	CreateTime time.Time
}

type PlanItem struct {
	Span  string
	Items []Item
}
