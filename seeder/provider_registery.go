package seeder

import (
	"github.com/nelsonsaake/go/register"
)

var Registry = register.NewRegistry[Seeder]()

func Get(name string) (Seeder, bool) {
	return Registry.Get(name)
}
