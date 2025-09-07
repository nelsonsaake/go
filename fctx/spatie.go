package fctx

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nelsonsaake/go/spatie/spatie"
)

// Use facade methods directly, no need for CheckFunc or singleton
type checkFunc func(userId string, args ...string) (bool, error)

func check(c *fiber.Ctx, args []string, f checkFunc) (bool, error) {
	die := func(f string, a ...any) (bool, error) {
		return false, fmt.Errorf(f, a...)
	}

	userId, err := UserID(c)
	if err != nil {
		return die("error getting user id: %w", err)
	}

	ok, err := f(userId, args...)
	if err != nil {
		return die("error checking user's ability: %w", err)
	}

	return ok, nil
}

func Is(c *fiber.Ctx, roles ...string) (bool, error) {
	return check(c, roles, spatie.Is)
}

func IsAny(c *fiber.Ctx, roles ...string) (bool, error) {
	return check(c, roles, spatie.IsAny)
}

func Can(c *fiber.Ctx, ps ...string) (bool, error) {
	return check(c, ps, spatie.Can)
}

func CanAny(c *fiber.Ctx, ps ...string) (bool, error) {
	return check(c, ps, spatie.CanAny)
}
