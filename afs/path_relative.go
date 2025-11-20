package afs

import (
	"fmt"
	"strings"
)

func RelativePath(fullPath string) (string, error) {

	base, err := Root()
	if err != nil {
		return "", fmt.Errorf("error getting root path: %w", err)
	}

	uri := strings.TrimPrefix(
		CleanPath(fullPath),
		CleanPath(base),
	)

	return uri, nil
}
