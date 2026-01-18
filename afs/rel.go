package afs

import (
	"strings"
)

func Rel(fullPath string) string {

	uri := strings.TrimPrefix(
		Clean(fullPath),
		Clean(root),
	)

	return uri
}
