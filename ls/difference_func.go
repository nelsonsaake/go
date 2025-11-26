package ls

// DifferenceFunc returns a new slice containing elements from slice a
// that are not present in slice b, determined by the provided equality function.
//
//	DifferenceFunc(usersA, usersB, func(x, y User) bool {
//		return x.ID == y.ID
//	})
func DifferenceFunc[T any](a, b []T, equal func(T, T) bool) []T {
	var result []T

	for _, v := range a {
		found := false
		for _, w := range b {
			if equal(v, w) {
				found = true
				break
			}
		}
		if !found {
			result = append(result, v)
		}
	}
	return result
}
