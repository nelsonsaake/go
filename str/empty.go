package str

import "strings"

func Empty(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}
