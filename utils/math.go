package utils

import (
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min >= max {
		return max
	}
	return rand.Intn(max-min) + min
}
