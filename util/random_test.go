package util

import (
	"fmt"
	"testing"
)

func TestRestaurant(t *testing.T) {
	m := make(map[int]int)
	for range 100000 {
		restaurant := RandRestaurant(3)
		m[restaurant]++
	}
	fmt.Println(m)
}
