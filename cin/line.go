package cin

import (
	"strings"
)

func Line(line *string) error {
	v, err := reader.ReadString('\n')
	v = strings.TrimSpace(v)
	*line = v
	return err
}
