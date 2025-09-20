package passport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonsaake/go/passport/middleware"
)

func newMiddleware(cfg TokenConfig, fctx *Fctx) fiber.Handler {
	return middleware.New(cfg, fctx)
}

func Middlewares() []fiber.Handler {

	res := []fiber.Handler{}

	for _, cfg := range cfgInstance.TokenConfigs {

		middleware := newMiddleware(cfg, getFctx())
		res = append(res, middleware)
	}

	return res
}
