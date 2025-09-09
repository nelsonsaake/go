package serr

import (
	"net/http"
)

func SendErrorWithStatus(status int) *ServerError {
	return New().WithStatus(status).WithMessage(http.StatusText(status))
}

func SendError() *ServerError {
	return SendErrorWithStatus(http.StatusUnprocessableEntity)
}

func BadRequestError() *ServerError {
	return SendErrorWithStatus(http.StatusBadRequest)
}

func UnauthorizedError() *ServerError {
	return SendErrorWithStatus(http.StatusUnauthorized)
}

func ForbiddenError() *ServerError {
	return SendErrorWithStatus(http.StatusForbidden)
}

func NotFoundError() *ServerError {
	return SendErrorWithStatus(http.StatusNotFound)
}

func MethodNotAllowedError() *ServerError {
	return SendErrorWithStatus(http.StatusMethodNotAllowed)
}

func ConflictError() *ServerError {
	return SendErrorWithStatus(http.StatusConflict)
}

func UnprocessableEntityError() *ServerError {
	return SendErrorWithStatus(http.StatusUnprocessableEntity)
}

func TooManyRequestsError() *ServerError {
	return SendErrorWithStatus(http.StatusTooManyRequests)
}

func InternalServerError() *ServerError {
	return SendErrorWithStatus(http.StatusInternalServerError)
}

func BadGatewayError() *ServerError {
	return SendErrorWithStatus(http.StatusBadGateway)
}

func ServiceUnavailableError() *ServerError {
	return SendErrorWithStatus(http.StatusServiceUnavailable)
}

func GatewayTimeoutError() *ServerError {
	return SendErrorWithStatus(http.StatusGatewayTimeout)
}
