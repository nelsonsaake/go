package rand

import "math/rand"

func Element[T any](ls ...T) T {

	if len(ls) == 0 {
		panic("cannot pick a random element from an empty list")
	}

	var (
		min = 0
		max = len(ls) - 1
	)

	i := rand.Intn(max-min+1) + min

	return ls[i]
}
