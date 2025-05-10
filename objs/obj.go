package objs

import (
	"iter"
	"maps"
)

type Obj struct {
	Data map[string]any
}

func New() *Obj {
	return &Obj{map[string]any{}}
}

func (m Obj) Keys() iter.Seq[string] {
	return maps.Keys(m.Data)
}

func (m Obj) Set(key string, value any) {
	Set(m.Data, key, value)
}

func (m Obj) Get(key string) any {
	return Get(m.Data, key)
}

// G: get any
func (m Obj) G(key string) any {
	return Get(m.Data, key)
}

func (m Obj) GetString(key string) string {
	return GetString(m.Data, key)
}

// GS: get string
func (m Obj) GS(key string) string {
	return GetString(m.Data, key)
}

func (m Obj) GetInt(key string) int {
	return GetInt(m.Data, key)
}

// GI: get int
func (m Obj) GI(key string) int {
	return GetInt(m.Data, key)
}

func (m Obj) GetFloat64(key string) float64 {
	return GetFloat64(m.Data, key)
}

// GF: get float64
func (m Obj) GF(key string) float64 {
	return GetFloat64(m.Data, key)
}

func (m Obj) GetBool(key string) bool {
	return GetBool(m.Data, key)
}

// GB: get bool
func (m Obj) GB(key string) bool {
	return GetBool(m.Data, key)
}

func (m Obj) GetMap(key string) map[string]any {
	return GetMap(m.Data, key)
}

// GM: get map[string]any
func (m Obj) GM(key string) map[string]any {
	return GetMap(m.Data, key)
}

func (m Obj) GetStringMap(key string) map[string]string {
	return GetStringMap(m.Data, key)
}

// GSM: get string map
func (m Obj) GSM(key string) map[string]string {
	return GetStringMap(m.Data, key)
}

func (m Obj) GetSlice(key string) []any {
	return GetSlice(m.Data, key)
}

// GSL: get slice []any
func (m Obj) GSL(key string) []any {
	return GetSlice(m.Data, key)
}

func (m Obj) GetStringSlice(key string) []string {
	return GetStringSlice(m.Data, key)
}

// GSS: get string slice
func (m Obj) GSS(key string) []string {
	return GetStringSlice(m.Data, key)
}

func (m Obj) GetObj(key string) *Obj {
	return GetObj(m.Data, key)
}

// GO: get Obj
func (m Obj) GO(key string) *Obj {
	return GetObj(m.Data, key)
}

func (m Obj) Cast(to any) error {
	return cast(m.Data, to)
}

func (m Obj) Delete(key string) any {
	return Delete(m.Data, key)
}

func (m Obj) String() string {
	return String(m.Data)
}
