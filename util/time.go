package util

import (
	"memo/constant"
	"time"
)

func LatePan(span string, done time.Time) time.Time {
	var update time.Time
	switch span {
	case constant.Day:
		update = time.Date(done.Year(), done.Month(), done.Day()+1, 0, 0, 0, 0, time.Local)
	case constant.Week:
		update = time.Date(done.Year(), done.Month(), done.Day()+8-int(done.Weekday()), 0, 0, 0, 0, time.Local)
	case constant.Month:
		update = time.Date(done.Year(), done.Month()+1, 1, 0, 0, 0, 0, time.Local)
	}
	return update
}
