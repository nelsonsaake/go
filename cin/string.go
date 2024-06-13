package cin

import (
	"strings"
)

func Line() (string, error) {
	v, err := reader.ReadString('\n')
	v = strings.TrimSpace(v)
	return v, err
}
