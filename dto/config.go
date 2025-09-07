package dto

import (
	"github.com/pinzolo/casee"
)

type Config struct {
	InclusionStrategy `json:"inclusion_strategy"`
	Case              `json:"case"`
}

func transformKey(k string) string {
	switch getKeyCase() {
	case PascalCase:
		k = casee.ToPascalCase(k)
	case CamelCase:
		k = casee.ToCamelCase(k)
	case SnakeCase:
		k = casee.ToSnakeCase(k)
	case ChainCase, KebabCase:
		k = casee.ToChainCase(k)
	}
	return k
}
