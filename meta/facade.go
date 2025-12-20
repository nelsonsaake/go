package meta

import "github.com/nelsonsaake/go/objs"

func Get(key string) any {
	return obj.Get(key)
}

func GetString(key string) string {
	return obj.GetString(key)
}

func GetInt(key string) int {
	return obj.GetInt(key)
}

func GetBool(key string) bool {
	return obj.GetBool(key)
}

func GetFloat64(key string) float64 {
	return obj.GetFloat64(key)
}

func GetMap(key string) map[string]any {
	return obj.GetMap(key)
}

func GetStringMap(key string) map[string]string {
	return obj.GetStringMap(key)
}

func GetSlice(key string) []any {
	return obj.GetSlice(key)
}

func GetStringSlice(key string) []string {
	return obj.GetStringSlice(key)
}

func GetObj(key string) *objs.Obj {
	return obj.GetObj(key)
}

func Set(key string, value any) {
	obj.Set(key, value)
}
