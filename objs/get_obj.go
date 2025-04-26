package objs

func GetObj(m map[string]any, k string) *Obj {
	v, ok := Get(m, k).(map[string]any)
	if !ok {
		return nil
	}
	return &Obj{v}
}
