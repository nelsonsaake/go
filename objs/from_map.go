package objs

func FromMap(v map[string]any) *Obj {
	return &Obj{v}
}
