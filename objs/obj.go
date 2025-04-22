package objs

type Map struct {
	Data map[string]any
}

func New() *Map {
	return &Map{map[string]any{}}
}

func FromMap(v map[string]any) *Map {
	return &Map{v}
}

func (m Map) Set(key string, value any) {
	Set(m.Data, key, value)
}

func (m Map) Get(key string) any {
	return Get(m.Data, key)
}

func (m Map) GetString(key string) string {
	return GetString(m.Data, key)
}

func (m Map) GetInt(key string) int {
	return GetInt(m.Data, key)
}

func (m Map) GetFloat64(key string) float64 {
	return GetFloat64(m.Data, key)
}

func (m Map) GetBool(key string) bool {
	return GetBool(m.Data, key)
}

func (m Map) GetMap(key string) map[string]any {
	return GetMap(m.Data, key)
}

func (m Map) GetStringMap(key string) map[string]string {
	return GetStringMap(m.Data, key)
}

func (m Map) GetSlice(key string) []any {
	return GetSlice(m.Data, key)
}

func (m Map) GetStringSlice(key string) []string {
	return GetStringSlice(m.Data, key)
}

func (m Map) Delete(key string) any {
	return Delete(m.Data, key)
}

func (m Map) String() string {
	return String(m.Data)
}
