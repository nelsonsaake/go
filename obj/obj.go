package obj

import "strings"

type Json map[string]any

func New(src map[string]any) Json {
	return src
}

func (src Json) Get(key string) any {

	var res any = src

	keys := strings.Split(key, ".")
	for _, key := range keys {

		data, ok := res.(map[string]any)
		if !ok {
			return nil
		}

		res = data[key]
	}

	return res
}

func (src Json) GetString(key string) string {
	return src.Get(key).(string)
}

func (src Json) GetInt(key string) int {
	return src.Get(key).(int)
}

func (src Json) GetFloat64(key string) float64 {
	return src.Get(key).(float64)
}

func (src Json) GetBool(key string) bool {
	return src.Get(key).(bool)
}
