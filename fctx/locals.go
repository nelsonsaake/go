package fctx

import (
	"github.com/nelsonsaake/go/serr"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrUnauthorized = serr.Message("You're not authorized to perform this action.").WithStatus(401)
)

func User(c *fiber.Ctx) (any, error) {

	user, ok := c.Locals("user").(any)
	if !ok {
		return nil, ErrUnauthorized.WithError("user not set")
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

func Auth(c *fiber.Ctx) (any, error) {

	auth, ok := c.Locals("auth").(any)
	if !ok {
		return nil, ErrUnauthorized.WithError("auth not set")
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
