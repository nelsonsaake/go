package objs

func GetMap(m map[string]any, k string) map[string]any {
	v, ok := Get(m, k).(map[string]any)
	if !ok {
		return nil
	}
	return v
}

func GetStringMap(m map[string]any, k string) map[string]string {

	var (
		from = Get(m, k)
		to   = map[string]string{}
	)

	if from == nil {
		return nil
	}

	err := cast(from, &to)
	if err != nil {
		return nil
	}

	return to
}

func GetSlice(m map[string]any, k string) []any {
	v, ok := Get(m, k).([]any)
	if !ok {
		return nil
	}
	return v
}

func GetStringSlice(m map[string]any, k string) []string {
	var (
		from = Get(m, k)
		to   = []string{}
	)

	err := cast(from, &to)
	if err != nil {
		return nil
	}

	return to
}
