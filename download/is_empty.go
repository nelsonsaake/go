package download

import "strings"

func isEmpty(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}
