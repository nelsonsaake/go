package arr

import "strings"

func IsEmpty(v string) bool {
	v = strings.TrimSpace(v)
	return len(v) == 0
}
