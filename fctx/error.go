package fctx

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, err error) error {
	log.Println(err)
	return err
}
