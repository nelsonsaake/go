package obj

import "strings"

type JsonMap map[string]any

func New(src map[string]any) JsonMap {
	return src
}

func (src JsonMap) Get(key string) any {

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

func (src JsonMap) GetString(key string) string {
	return src.Get(key).(string)
}

func (src JsonMap) GetInt(key string) int {
	return src.Get(key).(int)
}

func (src JsonMap) GetFloat64(key string) float64 {
	return src.Get(key).(float64)
}

func (src JsonMap) GetBool(key string) bool {
	return src.Get(key).(bool)
}
