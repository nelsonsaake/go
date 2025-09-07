package dto

var dtoRegister = map[string]DTO{}

func RegisterDTO(name string, dto DTO) {
	dtoRegister[name] = dto
}

func GetDTO(name string) DTO {
	dto, ok := dtoRegister[name]
	if ok {
		return dto
	}
	return DTO{}
}

var funcRegister = map[string]TransformValueFunc{}

func RegisterFunc(name string, f TransformValueFunc) {
	funcRegister[name] = f
}

func GetFunc(name string) TransformValueFunc {
	f, ok := funcRegister[name]
	if ok {
		return f
	}
	return func(name any) (any, error) { return nil, nil }
}

var calcValueRegister = map[string]CalcValueFunc{}

func RegisterCalcValueFunc(name string, f CalcValueFunc) {
	calcValueRegister[name] = f
}

func GetCalcValueFunc(v string) CalcValueFunc {
	f, ok := calcValueRegister[v]
	if ok {
		return f
	}
	return func(dto DTO) (any, error) { return nil, nil }
}
