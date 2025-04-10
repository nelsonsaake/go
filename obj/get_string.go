package obj

import "fmt"

func GetString(m map[string]any, k string) string {
	return fmt.Sprintf("%v", Get(m, k))
}
