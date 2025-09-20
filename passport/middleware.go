package passport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonsaake/go/passport/middleware"
)

func newMiddleware(cfg TokenConfig, fctx *Fctx) fiber.Handler {
	return middleware.New(cfg, fctx)
}

func Middleware(name string) fiber.Handler {

	cfg, exists := cfgInstance.Tokens[name]
	if !exists {
		return nil
	}

	return newMiddleware(cfg, getFctx())
}

func Middlewares() []fiber.Handler {

	res := []fiber.Handler{}

	for _, cfg := range cfgInstance.Tokens {

		middleware := newMiddleware(cfg, getFctx())
		res = append(res, middleware)
	}

	return res
}

func MiddlewareAuth() fiber.Handler {
	return Middleware("access")
}

func MiddlewareAuthRefresh() fiber.Handler {
	return Middleware("refresh")
}

func MiddlewareAuthReset() fiber.Handler {
	return Middleware("reset")
}
