package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonsaake/go/passport/models"
)

type Fctx interface {
	GetAuth(*fiber.Ctx) (*models.Auth, error)
	GetUser(*fiber.Ctx) (models.User, error)
}
