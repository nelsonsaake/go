package objs

func GetMap(m map[string]any, k string) map[string]any {
	v, ok := Get(m, k).(map[string]any)
	if !ok {
		return nil
	}
	return v
}

func GetSlice(m map[string]any, k string) []any {
	v, ok := Get(m, k).([]any)
	if !ok {
		return nil
	}
	return v
}

func GetStringSlice(m map[string]any, k string) []string {
	v, ok := Get(m, k).([]string)
	if !ok {
		return nil
	}
	return v
}
