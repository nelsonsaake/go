package fctx

import (
	"github.com/nelsonsaake/go/serr"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrUnauthorized = serr.Message("You're not authorized to perform this action.").WithStatus(401)
)

func User[T any](c *fiber.Ctx) (T, error) {

	user, ok := c.Locals("user").(T)
	if !ok {
		return *new(T), ErrUnauthorized.WithError("user not set")
	}

	return user, nil
}

func UserID(c *fiber.Ctx) (string, error) {

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return "", ErrUnauthorized.WithError("user id not set")
	}

	return userID, nil
}

func Auth[T any](c *fiber.Ctx) (T, error) {

	auth, ok := c.Locals("auth").(T)
	if !ok {
		return *new(T), ErrUnauthorized.WithError("auth not set")
	}

	return auth, nil
}

func AuthID(c *fiber.Ctx) (string, error) {

	userID, ok := c.Locals("authID").(string)
	if !ok {
		return "", ErrUnauthorized.WithError("auth id not set")
	}

	return userID, nil
}

func IsAuthorized(c *fiber.Ctx) bool {

	isAuthorized, ok := c.Locals("isAuthorized").(bool)
	return isAuthorized && ok
}
