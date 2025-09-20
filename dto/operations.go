package dto

type Operation interface {
	transformValue | transformType | includeValue | excludeValue | calcValue
}

// --- value transformation

type TransformValueFunc func(any) (any, error)

type transformValue struct {
	name string
}

func TransformValue(name string) transformValue {
	return transformValue{name}
}

// --- dto transformation type

type transformType struct {
	dtoname string
}

func TransformType(name string) transformType {
	return transformType{name}
}

// --- inclusion strategy

type includeValue struct{}

func Include() includeValue {
	return includeValue{}
}

// --- exclusion strategy

type excludeValue struct{}

func Exclude() excludeValue {
	return excludeValue{}
}

// ---

type CalcValueFunc func(dto DTO) (any, error)

type calcValue struct {
	name string
}

func CalcValue() excludeValue {
	return excludeValue{}
}
