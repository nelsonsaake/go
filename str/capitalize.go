package str

import (
	"fmt"
	"strings"
)

func _capitalize(v string) string {
	if Empty(v) {
		return v
	}
	fl := string(v[0])
	fl = strings.ToUpper(fl)
	return fmt.Sprintf("%s%s", fl, v[1:])
}

func Capitalize(v string) string {
	strs := strings.Split(v, " ")
	for i, str := range strs {
		if len(str) <= 3 {
			strs[i] = strings.ToLower(str)
			continue
		}
		if !Empty(str) {
			strs[i] = _capitalize(str)
		}
	}
	return strings.Join(strs, " ")
}
