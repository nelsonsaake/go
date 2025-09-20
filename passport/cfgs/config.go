package cfgs

import (
	"errors"
)

type Token struct {
	Name          string `json:"name"`
	SigningMethod string `json:"signing_method"`
	SigningKey    string `json:"signing_key"`
	TTL           int    `json:"ttl"`
}

func New(name string) (Token, error) {

	config, ok := Get(name)
	if ok {
		return config, nil
	}

	return config, errors.New("no config found for token: " + name)
}

func (t Token) GetName() string {
	return t.Name
}

func (t Token) GetSigningKey() []byte {
	return []byte(t.SigningKey)
}

func (t Token) GetTTL() int {
	return t.TTL
}

func (t Token) GetSigningMethod() string {
	return t.SigningMethod
}
