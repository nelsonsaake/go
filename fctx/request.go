package fctx

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Search       string   `query:"search" json:"search"`
	With         []string `query:"with" json:"with"`
	NoPagination bool     `query:"no_pagination" json:"no_pagination"`
	Page         int      `query:"page" json:"page"`
	PageSize     int      `query:"pageSize" json:"pageSize"`
}

type Paginated[T any] struct {
	Request
	Data       []T `json:"data"`
	Total      int64
	TotalPages int64
}

var zeroQuery = Request{
	Page:     1,
	PageSize: 20,
}

func GetQuery(c *fiber.Ctx) (Request, error) {

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

func (r *Request) JMap() map[string]any {
	return map[string]any{
		"search":        r.Search,
		"with":          r.With,
		"no_pagination": r.NoPagination,
		"page":          r.Page,
		"pageSize":      r.PageSize,
	}
}
