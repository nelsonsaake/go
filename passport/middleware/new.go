package middleware

import (
	"github.com/gofiber/fiber/v2"
	jware "github.com/gofiber/jwt/v3"
)

type Builder struct {
	Fctx
	config TokenConfig
}

func NewBuilder(fctx Fctx) *Builder {
	return &Builder{Fctx: fctx}
}

func (b *Builder) WithConfig(config TokenConfig) *Builder {
	b.config = config
	return b
}

func (b *Builder) build() fiber.Handler {

	middleware := jware.New(
		jware.Config{
			SigningKey:     b.config.GetSigningKey(),
			SigningMethod:  b.config.GetSigningMethod(),
			ErrorHandler:   b.onError,
			SuccessHandler: b.onSuccess,
		},
	)

	return middleware
}

func New(config TokenConfig, fctx Fctx) fiber.Handler {
	return NewBuilder(fctx).WithConfig(config).build()
}
