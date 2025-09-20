package middleware

type TokenConfig interface {
	GetName() string
	GetSigningKey() []byte
	GetTTL() int
	GetSigningMethod() string
}
