package vtp

import "reflect"

func IsString(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.String
}

func IsArray(v interface{}) bool {
	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}
