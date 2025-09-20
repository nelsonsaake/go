package fctx

import "github.com/gofiber/fiber/v2"

type Query struct {
	Search    string   `query:"search"`
	With      []string `query:"with"`
	Pagiation bool     `query:"pagination"`
	Page      int      `query:"page"`
	PageSize  int      `query:"pageSize"`
}

func GetQuery(c *fiber.Ctx) *Query {

	var q Query
	c.QueryParser(q)

	return &q
}
