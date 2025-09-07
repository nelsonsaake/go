package fctx

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonsaake/go/serr"
)

func SendErrorWithStatus(c *fiber.Ctx, status int, err error) error {

	err = serr.From(err)

	// log.Printf("[error]%w", err)

	return SendJSON(c.Status(status), err)
}

func SendError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrUnprocessableEntity.Code, err)
}

func BadRequestError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrBadRequest.Code, err)
}

func UnauthorizedError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrUnauthorized.Code, err)
}

func ForbiddenError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrForbidden.Code, err)
}

func NotFoundError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrNotFound.Code, err)
}

func MethodNotAllowedError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrMethodNotAllowed.Code, err)
}

func ConflictError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrConflict.Code, err)
}

func UnprocessableEntityError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrUnprocessableEntity.Code, err)
}

func TooManyRequestsError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrTooManyRequests.Code, err)
}

func InternalServerError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrInternalServerError.Code, err)
}

func BadGatewayError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrBadGateway.Code, err)
}

func ServiceUnavailableError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrServiceUnavailable.Code, err)
}

func GatewayTimeoutError(c *fiber.Ctx, err error) error {
	return SendErrorWithStatus(c, fiber.ErrGatewayTimeout.Code, err)
}
