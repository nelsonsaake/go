package vld

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func (v Validator) Validate(data any) *ValidatorErrors {

	err := v.validator.Struct(data) // err from underlying validator
	if err == nil {
		return nil
	}

	res := map[string][]string{}
	for _, err := range err.(validator.ValidationErrors) {
		if err != nil {

			v := interprete(err)
			k := getJSONTag(data, err)

			ls, ok := res[k]
			if !ok {
				ls = []string{}
			}

			ls = append(ls, v)

			res[k] = ls
		}
	}

	if len(res) > 0 {
		return &ValidatorErrors{res}
	}

	return nil
}

func New() *Validator {

	myValidator := &Validator{
		validator: validator.New(),
	}

	// Custom struct validation tag format
	myValidator.validator.RegisterValidation("teener", func(fl validator.FieldLevel) bool {
		// User.Age needs to fit our needs, 12-18 years old.
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})

	// Custom struct validation tag format
	// myValidator.validator.RegisterValidation("same", func(fl validator.FieldLevel) bool {
	// 	return fl.Field().Interface() == fl.Param()
	// })

	return myValidator
}
