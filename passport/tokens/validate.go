package tokens

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nelsonsaake/go/passport/cfgs"
)

func Validate(name, value string) (jwt.MapClaims, error) {

	resetTokenConfig, ok := cfgs.Get(name)
	if !ok {
		return nil, fmt.Errorf("error get %s token config", name)
	}

	token, err := jwt.ParseWithClaims(value, jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return resetTokenConfig.GetSigningKey(), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid and cast claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
