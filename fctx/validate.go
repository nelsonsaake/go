package fctx

import (
	"github.com/nelsonsaake/go/serr"
	"github.com/nelsonsaake/go/vld"
)

func Validate(data any) *serr.ServerError {

	err := vld.Validate(data)
	if err == nil || len(err.Errors) == 0 {
		return nil
	}

	return serr.New().
		WithMessage("The given data was invalid.").
		WithErrors(err.Errors)
}
