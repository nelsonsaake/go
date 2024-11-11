package rand 

import "math/rand"

func Bool() bool {
	return rand.Float32() >= 0.5
}