package xlscell

import (
	"errors"
	"strings"
)

// SplitRange: into component cells
func SplitRange(r string) (string, string, error) {
	cs := strings.Split(r, ":")

	start, end := "", ""

	switch len(cs) {
	case 1:
		start, end = cs[0], cs[0]
	case 2:
		start, end = cs[0], cs[1]
	default:
		start, end = "", ""
	}

	if !IsCell(start) || !IsCell(end) {
		return "", "", errors.New("is not a valid cell range")
	}

	return start, end, nil
}
