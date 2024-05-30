package str

func Contains(arr []string, e string) bool {

	for _, elem := range arr {
		if elem == e {
			return true
		}
	}

	return false
}
