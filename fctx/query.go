package fctx

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Query struct {
	Search      string   `query:"search"`
	With        []string `query:"with"`
	NoPagiation bool     `query:"no_pagination"`
	Page        int      `query:"page"`
	PageSize    int      `query:"pageSize"`
}

var zeroQuery = Query{
	Page:     1,
	PageSize: 20,
}

func GetQuery(c *fiber.Ctx) (Query, error) {

	var q = zeroQuery

	err := c.QueryParser(&q)
	if err != nil {
		return zeroQuery, fmt.Errorf("failed to parse query params: %w", err)
	}

	if q.Page == 0 {
		q.Page = 1
	}

	if q.PageSize == 0 {
		q.PageSize = 20
	}

	return q, nil
}
