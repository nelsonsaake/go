package objs

import "fmt"

func GetStrict[T any](m map[string]any, key string) T {
	switch any(*new(T)).(type) {
	case int:
		return any(GetInt(m, key)).(T)
	case string:
		return any(GetString(m, key)).(T)
	case bool:
		return any(GetBool(m, key)).(T)
	case float64:
		return any(GetFloat64(m, key)).(T)
	case []any:
		return any(GetSlice(m, key)).(T)
	case []string:
		return any(GetStringSlice(m, key)).(T)
	case map[string]any:
		return any(GetMap(m, key)).(T)
	case map[string]string:
		return any(GetStringMap(m, key)).(T)
	case *Obj:
		return any(GetObj(m, key)).(T)
	default:
		panic(fmt.Sprintf("GetStrict: unsupported type %T", *new(T)))
	}
}
