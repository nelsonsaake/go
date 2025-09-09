package serr

import (
	"net/http"
)

func SendErrorWithStatus(status int, f string, a ...any) *ServerError {
	return New().WithStatus(status).WithErrorf(f, a...)
}

func SendError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusUnprocessableEntity, f, a...)
}

func BadRequestError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusBadRequest, f, a...)
}

func UnauthorizedError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusUnauthorized, f, a...)
}

func ForbiddenError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusForbidden, f, a...)
}

func NotFoundError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusNotFound, f, a...)
}

func MethodNotAllowedError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusMethodNotAllowed, f, a...)
}

func ConflictError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusConflict, f, a...)
}

func UnprocessableEntityError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusUnprocessableEntity, f, a...)
}

func TooManyRequestsError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusTooManyRequests, f, a...)
}

func InternalServerError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusInternalServerError, f, a...)
}

func BadGatewayError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusBadGateway, f, a...)
}

func ServiceUnavailableError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusServiceUnavailable, f, a...)
}

func GatewayTimeoutError(f string, a ...any) *ServerError {
	return SendErrorWithStatus(http.StatusGatewayTimeout, f, a...)
}
