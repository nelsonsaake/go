package obj

import (
	"encoding/json"
	"strings"
)

func delete0(m map[string]any, k string) any {

	var mmm map[string]any
	var kkk string

	var v any = m
	var q = strings.Split(k, ".")

	for _, kk := range q {

		mm, ok := v.(map[string]any)
		if !ok {
			return nil
		}

		// save
		mmm = mm

		// reset scope
		v, kkk = mm[kk], kk
	}

	if mmm != nil {
		delete(mmm, kkk)
	}

	return v
}

func Delete(m map[string]any, k string) map[string]any {

	ks := strings.Split(k, ".")
	if len(ks) == 1 {
		delete(m, k)
		return m
	}

	// key in the current object
	k0 := ks[0]

	// nested key
	kx := strings.Join(ks[1:], ".")

	// get v0, is the current value of k from file
	v0, ok := m[k0]
	if !ok {
		return m
	}

	// convert v0 to json
	r0, err := json.Marshal(v0)
	if err != nil {
		delete(m, k0)
		return m
	}

	// convert the json, r0, to map
	mx := map[string]any{}
	err = json.Unmarshal(r0, &mx)
	if err != nil {
		delete(m, k0)
		return m
	}

	// set the value in the map,
	// if v0 is a map of any kind, this will add to it
	Delete(mx, kx)

	// we update parent map
	m[k0] = mx

	return m
}
