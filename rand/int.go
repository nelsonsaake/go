package rand

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func Int(max int) int {
	if max == 0 {
		return 0
	}
	return rng.Intn(max)
}

func IntBetween(min, max int) int {
	if min >= max {
		return min
	}
	return rng.Intn(max-min+1) + min
}
