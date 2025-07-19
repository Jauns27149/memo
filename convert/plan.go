package convert

import (
	"errors"
	"fmt"
	"memo/constant"
	"memo/model"
	"memo/util"
	"strings"
	"time"
)

func RowsToPlanItem(span string, list []string) (planItem model.PlanItem, flag bool, err error) {
	items := make([]model.Item, len(list))
	for i, v := range list {
		values := strings.Split(v, constant.Comma)
		if len(values) != 2 {
			err = errors.New("plan item format error")
			return
		}

		items[i] = model.Item{Content: values[0]}
		if value := strings.TrimSpace(values[1]); value != "" {
			t, err := time.Parse(time.DateTime, value)
			if err != nil {
				return planItem, flag, err
			}
			items[i].CreateTime = t
		}
		if !items[i].CreateTime.IsZero() && util.LatePan(span, items[i].CreateTime).Before(time.Now()) {
			items[i].CreateTime = time.Time{}
			flag = true
		}
	}

	planItem = model.PlanItem{Span: span, Items: items}
	return
}

func PlanItemsToRows(items []model.Item) []string {
	rows := make([]string, len(items))
	for i, v := range items {
		t := ""
		if !v.CreateTime.IsZero() {
			t = v.CreateTime.Format(time.DateTime)
		}
		rows[i] = fmt.Sprintf("%s,%s", v.Content, t)
	}
	return rows
}
