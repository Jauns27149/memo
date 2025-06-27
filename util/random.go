package util

import (
	"math/rand/v2"
	"time"
)

var r = rand.New(rand.NewPCG(
	uint64(time.Now().UnixNano()),
	uint64(time.Now().Nanosecond()),
))

func RandRestaurant(restaurants []string) string {
	n := r.IntN(len(restaurants))
	return restaurants[n]
}
