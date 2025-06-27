package util

import (
	"fmt"
	"testing"
)

func TestRestaurant(t *testing.T) {
	sample := []string{
		"肯德基",
		"麦当劳",
		"必胜客",
	}
	
	m := make(map[string]int, len(sample))
	for range 100000 {
		restaurant := RandRestaurant(sample)
		m[restaurant]++
	}
	fmt.Println(m)
}
