package passport

import "github.com/nelsonsaake/go/passport/cfgs"

type TokenConfig = cfgs.Token

func RegisterTokenConfig(name string, config TokenConfig) {
	cfgs.Register(name, config)
}

func GetToeknConfig(name string) (TokenConfig, bool) {
	return cfgs.Get(name)
}

func GetAllTokenConfigs() map[string]TokenConfig {
	return cfgs.GetAll()
}
