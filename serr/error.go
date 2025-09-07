package serr

import (
	"fmt"

	"github.com/nelsonsaake/go/pretty"
)

type ServerError struct {
	// http statu code
	Status int `json:"status"`

	// Message to be displayed to the user
	//
	// This allows us to clean up the error a bit,
	// so that it makes sense to the user and
	// give some us
	//
	// eg. `Something went wrong!`
	Message string `json:"message,omitempty"`

	// The actual error
	// eg. `Not enough balance in system account.`
	Err string `json:"error,omitempty"`

	// Usually for validation errors
	// where we map field names to error messages
	Errors any `json:"errors,omitempty"`
}

func (e *ServerError) WithStatus(v int) *ServerError {
	e.Status = v
	return e
}

func (e *ServerError) WithMessage(v ...any) *ServerError {
	e.Message = fmt.Sprint(v...)
	return e
}

func (e *ServerError) WithMessagef(f string, v ...any) *ServerError {
	e.Message = fmt.Sprintf(f, v...)
	return e
}

func (e *ServerError) WithError(v ...any) *ServerError {
	e.Err = fmt.Sprint(v...)
	return e
}

func (e *ServerError) WithErrorf(f string, v ...any) *ServerError {
	e.Err = fmt.Sprintf(f, v...)
	return e
}

func (e *ServerError) WrapErrorf(f string, v ...any) *ServerError {
	e.Err = fmt.Sprintf(f, append(v, e.Err)...)
	return e
}

func (e *ServerError) WithErrors(v any) *ServerError {
	e.Errors = v
	return e
}

func (e *ServerError) WithStatusMessage(status int, v ...any) *ServerError {
	if status == e.Status {
		e = e.WithMessage(v...)
	}
	return e
}

func (e *ServerError) Error() string {
	return pretty.JSON(e)
}
