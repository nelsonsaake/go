package afs

import (
	"fmt"
	"strings"

	"github.com/nelsonsaake/go/ufs"
)

func RelativePath(fullPath string) (string, error) {

	base, err := Root()
	if err != nil {
		return "", fmt.Errorf("error getting root path: %w", err)
	}

	uri := strings.TrimPrefix(
		ufs.CleanPath(fullPath),
		ufs.CleanPath(base),
	)

	return uri, nil
}
