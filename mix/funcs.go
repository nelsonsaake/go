package mix

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/nelsonsaake/go/tfm"
)

func funcs() template.FuncMap {

	fls := sprig.FuncMap()
	fls["Join"] = tfm.FuncJoin
	return fls
}
