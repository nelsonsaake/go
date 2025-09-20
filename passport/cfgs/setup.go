package cfgs

func Setup(configs map[string]Token) {
	for key, config := range configs {
		Register(key, config)
	}
}
