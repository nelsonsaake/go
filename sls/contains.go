package sls

import "slices"

func Contains[T any](s []T, match func(T) bool) bool {
	return slices.ContainsFunc(s, match)
}
