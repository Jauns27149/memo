package util

import (
	"fmt"
	"memo/constant"
	"testing"
	"time"
)

func TestLatePan(t *testing.T) {
	sample := []struct {
		span   string
		create time.Time
	}{
		{constant.Day, time.Now().Add(-24 * time.Hour)},
		{constant.Week, time.Now().Add(-24 * time.Hour)},
		{constant.Month, time.Now().Add(-24 * time.Hour)},
	}
	for _, v := range sample {
		fmt.Println(v.create)
		fmt.Println(LatePan(v.span, v.create))
		fmt.Println()
	}
}
