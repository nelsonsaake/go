package rand

import (
	"math/rand"
)

func Int(max int) int {
	if max == 0 {
		return 0
	}
	return rand.Intn(max)
}

func IntBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}
