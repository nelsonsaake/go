package objs

import "strings"

func Get(m map[string]any, k string) any {

	var res any = m

	keys := strings.Split(k, ".")

	for _, key := range keys {

		_res, ok := res.(map[string]any)
		if !ok {
			return nil
		}

		res = _res[key]
	}

	return res
}
