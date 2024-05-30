package gen

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Int(max int) int {
	if max == 0 {
		return 0
	}
	return rand.Intn(max)
}
