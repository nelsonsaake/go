package cfgs

import (
	"fmt"
)

// getck: get cache key
func getck(t, k string) string {
	return fmt.Sprintf("%s:%s", t, k)
}

func typeString[T any]() string {
	return fmt.Sprintf("%T", *new(T))
}
