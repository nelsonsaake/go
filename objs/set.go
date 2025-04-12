package objs

import (
	"encoding/json"
	"strings"
)

func Set(m map[string]any, k string, v any) map[string]any {

	ks := strings.Split(k, ".")
	if len(ks) == 1 {
		m[k] = v
		return m
	}

	// key in the current object
	k0 := ks[0]

	// nested key
	kx := strings.Join(ks[1:], ".")

	// get v0, is the current value of k from file
	v0, ok := m[k0]
	if !ok {
		v0 = map[string]any{}
	}

	// convert v0 to json
	r0, err := json.Marshal(v0)
	if err != nil {
		r0 = []byte{}
	}

	// convert the json, r0, to map
	mx := map[string]any{}
	json.Unmarshal(r0, &mx)

	// set the value in the map,
	// if v0 is a map of any kind, this will add to it
	Set(mx, kx, v)

	// we update parent map
	m[k0] = mx

	return m
}
