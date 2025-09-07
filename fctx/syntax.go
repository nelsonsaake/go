package fctx

import "github.com/gofiber/fiber/v2"

func IsSyntaxRequest(c *fiber.Ctx) bool {
	return c.Query("syntax") == "true"
}
