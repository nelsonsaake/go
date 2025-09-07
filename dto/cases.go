package dto

type Case string

const (
	PascalCase Case = "pascal_case"
	CamelCase  Case = "camel_case"
	SnakeCase  Case = "snake_case"
	ChainCase  Case = "chain_case"
	KebabCase  Case = "kebab_case"
)

var (
	keyCase = CamelCase
)

func SetKeyCase(v string) {
	keyCase = Case(v)
}

func getKeyCase() Case {
	return keyCase
}
