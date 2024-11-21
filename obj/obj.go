package obj

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Map struct {
	Data map[string]any
}

func New(v map[string]any) *Map {
	return &Map{v}
}

func (m Map) Get(key string) any {

	var res any = m.Data

	keys := strings.Split(key, ".")

	for _, key := range keys {

		_res, ok := res.(map[string]any)
		if !ok {
			return nil
		}

		res = _res[key]
	}

	return res
}

func (m Map) GetString(key string) string {
	return fmt.Sprintf("%v", m.Get(key))
}

func (m Map) GetInt(key string) int {
	i, _ := strconv.Atoi(m.GetString(key))
	return i
}

func (m Map) GetFloat64(key string) float64 {
	return m.Get(key).(float64)
}

func (m Map) GetBool(key string) bool {
	return m.Get(key).(bool)
}

func (m Map) Delete(key string) any {

	var prev map[string]any
	var k string

	// v = prev[k]
	var v any = m
	var q = strings.Split(key, ".")

	for _, key := range q {

		_v, ok := v.(map[string]any)
		if !ok {
			return nil
		}

		// save
		prev = _v

		// reset scope
		v, k = _v[key], key
	}

	if prev != nil {
		delete(prev, k)
	}

	return v
}

func (m Map) String() string {
	jsonBytes, err := json.MarshalIndent(m.Data, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
