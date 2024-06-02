package fctx

import "github.com/gofiber/fiber/v2"

func String(c *fiber.Ctx, s string) error {
	return c.SendString(s)
}
