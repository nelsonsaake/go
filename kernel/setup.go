package kernel

type Config struct {
	Name    string
	Verbose bool
}

func Setup(c Config) {
	name = c.Name
	SetVerbose(c.Verbose)
}
