package vld

func Validate(data any) *ValidatorErrors {
	return New().Validate(data)
}
