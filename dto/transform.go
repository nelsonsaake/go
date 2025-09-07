package dto

import (
	"reflect"
)

func Transform(data any, dtoname string) (any, error) {

	var (
		kind    = reflect.TypeOf(data).Kind()
		isSlice = kind == reflect.Slice
		isArray = kind == reflect.Array
	)

	if isSlice || isArray {
		return transformls(data, dtoname)
	} else {
		return transformv(data, dtoname)
	}
}
