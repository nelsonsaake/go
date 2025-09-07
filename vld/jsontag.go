package vld

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func getJSONTag(model any, fieldError validator.FieldError) string {
	t := reflect.TypeOf(model)

	// Handle pointer to struct
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Ensure it's a struct before proceeding
	if t.Kind() != reflect.Struct {
		return fieldError.Field() // Fallback to field name if not a struct
	}

	field, found := t.FieldByName(fieldError.StructField())
	if !found {
		return fieldError.Field() // Fallback to struct field name
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return fieldError.Field() // If no JSON tag, return the field name
	}

	return jsonTag
}
