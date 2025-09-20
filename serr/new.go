package serr

import (
	"net/http"
)

func New() *ServerError {
	return &ServerError{
		Status:  http.StatusInternalServerError,
		Message: "Something went wrong! Please try again later.",
	}
}

func Errorf(f string, v ...any) *ServerError {
	return New().WithErrorf(f, v...)
}

func NotFound() *ServerError {
	return New().WithStatus(http.StatusNotFound)
}

func Message(v ...any) *ServerError {
	return New().WithMessage(v...)
}
