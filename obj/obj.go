package obj

import "strings"

type Map map[string]any

func New(v map[string]any) Map {
	return v
}

func (m Map) Get(key string) any {

	var res any = m

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

func (m Map) GetString(key string) string {
	return m.Get(key).(string)
}

func (m Map) GetInt(key string) int {
	return m.Get(key).(int)
}

func (m Map) GetFloat64(key string) float64 {
	return m.Get(key).(float64)
}

func (m Map) GetBool(key string) bool {
	return m.Get(key).(bool)
}
