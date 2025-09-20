package passport

import "github.com/nelsonsaake/go/passport/cfgs"

type TokenConfig = cfgs.Token

func cfgsRegister(name string, config TokenConfig) {
	cfgs.Register(name, config)
}

func cfgsGet(name string) (TokenConfig, bool) {
	return cfgs.Get(name)
}

func cfgsGetAll() map[string]TokenConfig {
	return cfgs.GetAll()
}
