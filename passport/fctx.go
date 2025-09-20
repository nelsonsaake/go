package passport

import "github.com/nelsonsaake/go/passport/fctx"

type Fctx = fctx.Fctx

func getFctx() *Fctx {
	return fctx.New(getAuthRepo(), getUserRepo())
}
