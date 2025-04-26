package objs

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
