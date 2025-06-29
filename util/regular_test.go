package util

import (
	"fmt"
	"testing"
)

func TestGetItemNoTime(t *testing.T) {
	texts := [][2]string{
		{"3", "3"},
		{"1 2025-06-26 16:13:43", "1"},
	}

	for _, test := range texts {
		if result := GetItemNoTime(test[0]); result != test[1] {
			fmt.Printf("reuslt:%v\nanswer:%v\n", result, test[1])
		}
	}
}

func TestRestaurantHadTime(t *testing.T) {
	result := GetItemNoTime("1")
	if result != "1" {
		panic(result)
	}
}

func TestClearTime(t *testing.T) {
	result := ClearTime("sample 2025-06-26 16:13:43")
	if result != "sample" {
		println(result)
	}
}
