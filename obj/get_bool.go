package obj

func GetBool(m map[string]any, k string) bool {
	return Get(m, k).(bool)
}
