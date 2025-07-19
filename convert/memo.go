package convert

import (
	"errors"
	"fmt"
	"memo/constant"
	"memo/model"
	"strings"
	"time"
)

func MemosToRows(items []model.MemoItem) (rows []string) {
	rows = make([]string, len(items))
	for i, item := range items {
		rows[i] = fmt.Sprintf("%s,%s,%d",
			item.Content,
			item.CreateTime.Format(time.DateTime),
			boolToInt(item.Finished))
	}
	return
}

func RowsToMemo(rows []string) (item []model.MemoItem, err error) {
	for _, row := range rows {
		memo, err := RowToMemo(row)
		if err != nil {
			return nil, err
		}
		item = append(item, memo)
	}
	return
}

func RowToMemo(row string) (item model.MemoItem, err error) {
	row = strings.Trim(row, constant.Space)
	values := strings.Split(row, constant.Comma)
	if len(values) != 3 {
		return item, errors.New("invalid row: " + row)
	}
	createTime, err := time.Parse(time.DateTime, values[1])
	if err != nil {
		return item, err
	}
	var finished bool
	if values[2] == "0" {
		finished = false
	} else if values[2] == "1" {
		finished = true
	} else {
		return item, errors.New("finished field invalid row")
	}

	item = model.MemoItem{
		Item: model.Item{
			Content:    values[0],
			CreateTime: createTime,
		},
		Finished: finished,
	}
	return
}
