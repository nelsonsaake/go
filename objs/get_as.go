package objs

func GetAs(m map[string]any, k string, as any) {

	var (
		from = Get(m, k)
		to   = as
	)

	if from == nil {
		return
	}

	err := cast(from, to)
	if err != nil {
		return
	}
}
