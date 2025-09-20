package fctx

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nelsonsaake/go/passport/models"
)

type Fctx struct {
	Auths AuthRepo
	Users UserRepo
}

func New(Auths AuthRepo, Users UserRepo) *Fctx {
	return &Fctx{Auths: Auths, Users: Users}
}

func (f *Fctx) GetJwt(c *fiber.Ctx) (*jwt.Token, error) {

	tkn, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf("jwt is invalid or missing")
	}

	return tkn, nil
}

func (f *Fctx) GetJwtID(c *fiber.Ctx) (string, error) {

	die := func() (string, error) {
		return "", fmt.Errorf("jwt is invalid or missing")
	}

	tkn, err := f.GetJwt(c)
	if err != nil {
		return die()
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok {
		return die()
	}

	id, ok := claims["id"].(string)
	if !ok {
		return die()
	}

	return id, nil
}

func (f *Fctx) GetAuth(c *fiber.Ctx) (*models.Auth, error) {

	die := func(f string, err error) (*models.Auth, error) {
		return nil, fmt.Errorf(f, err)
	}

	id, err := f.GetJwtID(c)
	if err != nil {
		return die("error getting jwt: %w", err)
	}

	auth, err := f.Auths.Find(id)
	if err != nil {
		return die("error finding auth: %w", err)
	}
	if auth == nil {
		return die("jwt is invalid or missing: %v", nil)
	}

	return auth, nil
}

func (f *Fctx) GetUser(c *fiber.Ctx) (models.User, error) {

	die := func(f string, err error) (models.User, error) {
		return nil, fmt.Errorf(f, err)
	}

	auth, err := f.GetAuth(c)
	if err != nil {
		return die("error getting auth: %w", err)
	}

	user, err := f.Users.Find(auth.UserID)
	if err != nil {
		return die("error finding user: %w", err)
	}
	if user == nil {
		return die("jwt is invalid or missing: %v", nil)
	}

	return user, nil
}
