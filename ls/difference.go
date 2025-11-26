package ls

func Difference[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, v := range b {
		m[v] = true
	}

	var result []T
	for _, v := range a {
		if !m[v] {
			result = append(result, v)
		}
	}

	return result
}
