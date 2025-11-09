package dsc

import "strings"

func ID(id ...string) string {
	return strings.Join(id, "/")
}
