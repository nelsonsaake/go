package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// onSuccess: save auth credentials in context so we don't have fetch them every time
func (b *Builder) onSuccess(c *fiber.Ctx) error {

	// add auth to context

	auth, err := b.GetAuth(c)
	if err != nil {
		return fmt.Errorf("error getting auth: %w", err)
	}

	c.Locals("auth", auth)
	c.Locals("authID", auth.ID)

	// add user to context

	user, err := b.GetUser(c)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	c.Locals("user", user)
	c.Locals("userID", auth.UserID)

	// mark request as authorized

	c.Locals("isAuthorized", true)

	return c.Next()
}
