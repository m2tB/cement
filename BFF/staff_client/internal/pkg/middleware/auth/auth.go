package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type CusClaims struct {
	Mobile string
	Name   string
	jwt.RegisteredClaims
}

// CreateToken generate token
func CreateToken(c CusClaims, key string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("generate token failed" + err.Error())
	}
	return signedString, nil
}
