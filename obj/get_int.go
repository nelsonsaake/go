package obj

import "strconv"

func GetInt(m map[string]any, k string) int {
	i, _ := strconv.Atoi(GetString(m, k))
	return i
}
