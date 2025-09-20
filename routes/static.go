package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddStaticRoutes(app *fiber.App) {
	// Apply CORS globally
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS")

		// Handle preflight requests (OPTIONS)
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	// Static routes
	app.Static("/public", "./public")
	app.Static("/storage", "./storage")
	app.Static("/resources", "./resources")
}
