package strs

import "strings"

func IsEmpty(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}
