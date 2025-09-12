package cfgs

type SetupConfig struct {
	Dirs []string
}

func Setup(c SetupConfig) {
	Load(c.Dirs...)
}
