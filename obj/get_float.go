package obj

func GetFloat64(m map[string]any, k string) float64 {
	return Get(m, k).(float64)
}
