package objs

type Obj struct {
	Data map[string]any
}

func New() *Obj {
	return &Obj{map[string]any{}}
}

func FromMap(v map[string]any) *Obj {
	return &Obj{v}
}

func (m Obj) Set(key string, value any) {
	Set(m.Data, key, value)
}

func (m Obj) Get(key string) any {
	return Get(m.Data, key)
}

func (m Obj) GetString(key string) string {
	return GetString(m.Data, key)
}

func (m Obj) GetInt(key string) int {
	return GetInt(m.Data, key)
}

func (m Obj) GetFloat64(key string) float64 {
	return GetFloat64(m.Data, key)
}

func (m Obj) GetBool(key string) bool {
	return GetBool(m.Data, key)
}

func (m Obj) GetMap(key string) map[string]any {
	return GetMap(m.Data, key)
}

func (m Obj) GetStringMap(key string) map[string]string {
	return GetStringMap(m.Data, key)
}

func (m Obj) GetSlice(key string) []any {
	return GetSlice(m.Data, key)
}

func (m Obj) GetStringSlice(key string) []string {
	return GetStringSlice(m.Data, key)
}

func (m Obj) Delete(key string) any {
	return Delete(m.Data, key)
}

func (m Obj) String() string {
	return String(m.Data)
}
