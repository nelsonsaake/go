package middleware

import "time"

type TokenConfig interface {
	GetName() string
	GetSigningKey() []byte
	GetTTL() time.Duration
	GetSigningMethod() string
}
