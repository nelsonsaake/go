package cfgs

import (
	"fmt"
)

// makeCacheKey: get cache key
func makeCacheKey(t, k string) string {
	return fmt.Sprintf("%s:%s", t, k)
}

func typeString[T any]() string {
	return fmt.Sprintf("%T", *new(T))
}
