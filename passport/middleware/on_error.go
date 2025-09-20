package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nelsonsaake/go/serr"
)

// onError: inspect the error that occurred to give user, an actionable information
func (b *Builder) onError(c *fiber.Ctx, err error) error {

	if err == nil {
		return nil
	}

	var ve *jwt.ValidationError

	if !errors.As(err, &ve) {
		return serr.AuthError("error validating jwt: %v", err)
	}

	switch ve.Errors {
	case jwt.ValidationErrorExpired:
		return serr.JWTExpired("jwt expired: %v", err)
	case jwt.ValidationErrorSignatureInvalid:
		return serr.InvalidJWT("jwt signature invalid: %v", err)
	default:
		return serr.InvalidJWT("jwt invalid: %v", err)
	}
}
