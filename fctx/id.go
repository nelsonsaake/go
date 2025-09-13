package fctx

import "github.com/gofiber/fiber/v2"

func ID(c *fiber.Ctx) string {
	return c.Params("id")
}
