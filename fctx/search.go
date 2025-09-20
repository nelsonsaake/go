package fctx

import "github.com/gofiber/fiber/v2"

func Search(c *fiber.Ctx) string {
	return c.Query("search")
}
