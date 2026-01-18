package afs

import (
	"fmt"
	"strings"
)

func Rel(fullPath string) (string, error) {

	base, err := Root()
	if err != nil {
		return "", fmt.Errorf("error getting root path: %w", err)
	}

	uri := strings.TrimPrefix(
		Clean(fullPath),
		Clean(base),
	)

	return uri, nil
}
