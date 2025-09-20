package tokens

import (
	"fmt"
	"maps"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nelsonsaake/go/passport/cfgs"
)

func newTokenString[T string | []byte](
	id string,
	ttl time.Duration,
	signingKey T,
	signingMethod string,
	argClaims ...map[string]any,
) (string, error) {

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(ttl).Unix(),
	}

	for _, ac := range argClaims {
		maps.Copy(claims, ac)
	}

	tokenObject := jwt.NewWithClaims(
		jwt.GetSigningMethod(signingMethod),
		claims,
	)

	tokenString, err := tokenObject.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func New(name, id string, argClaims ...map[string]any) (string, error) {

	die := func(f string, a ...any) (string, error) {
		return "", fmt.Errorf(f, a...)
	}

	tcfg, ok := cfgs.Get(name)
	if !ok {
		return die("%v token config not available", name)
	}

	ttl := tcfg.GetTTL()
	if ttl == 0 {
		return die("%v token ttl not set", name)
	}

	signingKey := tcfg.GetSigningKey()
	if len(signingKey) == 0 {
		return die("%v token signing key not set", name)
	}

	return newTokenString(id, ttl, signingKey, tcfg.SigningMethod, argClaims...)
}
