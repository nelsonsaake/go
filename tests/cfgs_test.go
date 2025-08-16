package tests

import (
	"testing"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/cfgs"
)

func TestCfgs(t *testing.T) {
	c := cfgs.New(
		afs.Path("src/configs"),
	)
	t.Log(c.Get("dbs.model"))
	t.Log(c.Get("dbs.model.excluded"))
}
