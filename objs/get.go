package objs

import "strings"

func Get(m map[string]any, k string) any {

	if m == nil {
		return nil
	}

	if len(strings.TrimSpace(k)) == 0 {
		return nil
	}

	v, ok := m[k]
	if ok {
		return v
	}

	ks := strings.Split(k, ".")
	if len(ks) == 1 {
		return nil
	}

	// key in the current object
	k0 := ks[0]

	// nested key
	kk := strings.Join(ks[1:], ".")

	mm, ok := m[k0].(map[string]any)
	if !ok {
		return nil
	}

	return Get(mm, kk)
}
