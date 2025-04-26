package objs

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
