package repo

import "reflect"

// getTypeName returns the base type name of any value i,
// unwrapping pointers and slices.
func getTypeName(i interface{}) string {

	t := reflect.TypeOf(i)

	// Unwrap pointer or slice until you get the base type
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	return t.Name()
}
