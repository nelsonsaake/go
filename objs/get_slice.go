package objs

func GetSlice(m map[string]any, k string) []any {
	v, ok := Get(m, k).([]any)
	if !ok {
		return nil
	}
	return v
}
