package auth

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

type CusClaims struct {
	ID     int64
	Mobile string
	Name   string
	jwt2.RegisteredClaims
}

// GenerateToken generate token
func GenerateToken(c CusClaims, key string) (string, error) {
	claims := jwt2.NewWithClaims(jwt2.SigningMethodHS256, c)
	signedString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("generate token failed" + err.Error())
	}
	return signedString, nil
}

// GetClaimFormContext 从context上下文中获取claim信息
func GetClaimFormContext(ctx context.Context) jwt2.MapClaims {
	claims, _ := jwt.FromContext(ctx)
	return claims.(jwt2.MapClaims)
}
