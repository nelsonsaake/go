package tests

import (
	"testing"

	"github.com/nelsonsaake/go/cfgs"
)

func TestCfgs(t *testing.T) {
	t.Log(cfgs.Get("dbs/model"))
}
