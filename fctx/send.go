package fctx

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func SendString(c *fiber.Ctx, s string) error {

	return c.SendString(s)
}

func SendJSON(c *fiber.Ctx, data any) error {

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.SendString(string(jsonBytes))
}

func SendAs(c *fiber.Ctx, data any, dtoname string) error {

	data, err := Transform(data, dtoname)
	if err != nil {
		return UnprocessableEntityError(c, err)
	}

	return SendJSON(c, data)
}

func SendMessage(c *fiber.Ctx, message string) error {

	return send(c, Message(message))
}

func SendHtml(c *fiber.Ctx, html string) error {

	c.Set("Content-Type", "text/html")

	return c.Send([]byte(html))
}

func send(c *fiber.Ctx, fs ...func(*Response)) error {

	res := &Response{}
	for _, f := range fs {
		f(res)
	}

	return SendJSON(c, res)
}
