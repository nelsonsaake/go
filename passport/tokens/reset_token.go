package tokens

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nelsonsaake/go/passport/cfgs"
)

func NewResetToken(email string) (string, error) {

	return New("reset", email)
}

func VerifyResetToken(tokenString string) (jwt.MapClaims, error) {

	resetTokenConfig, ok := cfgs.Get("reset")
	if !ok {
		return nil, fmt.Errorf("error get reset token config")
	}

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
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

func IsValidResetToken(tokenString string, email string) bool {

	claims, err := VerifyResetToken(tokenString)
	if err != nil {
		return false
	}

	err = claims.Valid()
	if err != nil {
		return false
	}

	claimsEmail, ok := (claims["email"]).(string)
	if !ok {
		return false
	}

	return claimsEmail == email
}
