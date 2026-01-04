package do

import "reflect"

func MergeMaps(a, b any) any {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	if av.Kind() != reflect.Map || bv.Kind() != reflect.Map {
		panic("both arguments must be maps")
	}

	ak := av.Type().Key()
	bk := bv.Type().Key()

	// Rule: key types must match
	if ak != bk {
		panic("map key types do not match")
	}

	avt := av.Type().Elem()
	bvt := bv.Type().Elem()

	// Same key AND value type → preserve type
	if avt == bvt {
		out := reflect.MakeMap(av.Type())

		for _, k := range av.MapKeys() {
			out.SetMapIndex(k, av.MapIndex(k))
		}
		for _, k := range bv.MapKeys() {
			out.SetMapIndex(k, bv.MapIndex(k))
		}

		return out.Interface()
	}

	// Same key, different value type → map[K]any
	outType := reflect.MapOf(ak, reflect.TypeOf((*any)(nil)).Elem())
	out := reflect.MakeMap(outType)

	for _, k := range av.MapKeys() {
		out.SetMapIndex(k, av.MapIndex(k).Convert(outType.Elem()))
	}
	for _, k := range bv.MapKeys() {
		out.SetMapIndex(k, bv.MapIndex(k).Convert(outType.Elem()))
	}

	return out.Interface()
}
