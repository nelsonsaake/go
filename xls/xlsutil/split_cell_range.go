package xlsutil

import "strings"

func SplitCellRange(r string) (string, string) {
	cs := strings.Split(r, ":")

	switch len(cs) {
	case 1:
		return cs[0], cs[0]
	case 2:
		return cs[0], cs[1]
	default:
		return "", ""
	}
}
