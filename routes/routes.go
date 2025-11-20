package routes

import (
	"github.com/gofiber/fiber/v2"
)

type registerfunc func(*fiber.App)

var rfs = []registerfunc{}

func Register(r func(*fiber.App)) {
	rfs = append(rfs, r)
}

func Load(app *fiber.App) {
	for _, f := range rfs {
		f(app)
	}
}
