package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type LoginClaims struct {
	UserName string
	UserID interface{}
	jwt.StandardClaims
}

func GenerateToken(userName string, userID interface{}, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims {
		UserName: userName,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})

	return token.SignedString([]byte(ManagerConfig.SecretKey))
}