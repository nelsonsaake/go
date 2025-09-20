package cfgs

import "github.com/nelsonsaake/go/register"

var configs = register.New[Token]()

func Register(key string, config Token) {
	configs.Register(key, config)
}

func Get(key string) (Token, bool) {
	return configs.Get(key)
}

func GetAll() map[string]Token {
	return configs.GetAll()
}
