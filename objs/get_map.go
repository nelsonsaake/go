package objs

func GetMap(m map[string]any, k string) map[string]any {
	v, ok := Get(m, k).(map[string]any)
	if !ok {
		return nil
	}
	return v
}
