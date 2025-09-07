package fctx

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func BindBody(c *fiber.Ctx, out any) error {

	err := c.BodyParser(out)
	if err != nil {
		return fmt.Errorf("error parsing request body: %w", err)
	}

	verr := Validate(out)
	if verr != nil {
		return verr
	}

	return nil
}
