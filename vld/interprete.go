package vld

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// interprete: organize error in a user-friendly way like in laravel
func interprete(err validator.FieldError) string {
	field := toWords(err.Field())

	tr := map[string]string{
		"required": fmt.Sprintf("%s is required.", field),
		"email":    fmt.Sprintf("The %s provided is not a valid email.", field),
		"gte":      fmt.Sprintf("The %s must be greater than or equal to %s.", field, err.Param()),
		"lte":      fmt.Sprintf("The %s must be less than or equal to %s.", field, err.Param()),
		"eqfield":  fmt.Sprintf("The %s must be equal to %s.", field, err.Param()),
	}

	msg, exists := tr[err.Tag()]

	if !exists {
		// msg = fmt.Sprintf("The %s is not valid.", field)
		msg = err.Error() // default error message from validator
	}

	return toSentenceCase(msg)
}
