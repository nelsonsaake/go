package xlscell

import "math"

func minmax(a, b int) (int, int) {
	a64, b64 := float64(a), float64(b)
	a64, b64 = math.Min(a64, b64), math.Max(a64, b64)
	return int(a64), int(b64)
}
