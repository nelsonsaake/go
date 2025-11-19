package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var registry *fiber.App

func init() {
	registry = fiber.New()
}

func Register(r func(*fiber.App)) {
	r(registry)
}

func Load(app *fiber.App) {
	app.Use(logger.New())
	app.Mount("/", registry)
}
