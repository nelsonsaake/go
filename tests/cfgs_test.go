package tests

import (
	"fmt"
	"testing"

	"github.com/nelsonsaake/go/cfgs"
	"github.com/nelsonsaake/go/pty"
)

func TestCfgs(t *testing.T) {
	c := cfgs.New(
		"src/configs",
	)

	ln := func(arg ...int) {
		ln := ""

		var count = 0

		if len(arg) == 0 {
			count = 2
		}

		for i := range arg {
			count += i
		}

		for range count {
			ln += "\n"
		}

		fmt.Print(ln)
	}

	t.Log("dbs.model")
	ln(1)

	t.Log(pty.JSON(c.Get("dbs.model")))
	ln()

	t.Log("dbs.model.excluded")
	ln(1)

	t.Log(pty.JSON(c.Get("dbs.model.excluded")))
	ln()

	t.Log("meta")
	ln(1)

	t.Log(pty.JSON(c.Get("meta")))
	ln()
}
