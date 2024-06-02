package fctx

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func JSON(c *fiber.Ctx, data any) error {

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.SendString(string(jsonBytes))
}
