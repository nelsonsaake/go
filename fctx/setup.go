package fctx

type Config struct {
	TransformFunc DTOFunc
}

func Setup(c Config) {
	transform = c.TransformFunc
}
