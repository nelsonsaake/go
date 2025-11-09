package ls

func Contains[T any](s []T, match func(T) bool) bool {
	for _, v := range s {
		if match(v) {
			return true
		}
	}
	return false
}
