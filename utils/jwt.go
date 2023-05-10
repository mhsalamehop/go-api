package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey = []byte(os.Getenv("jwt_secret"))
)

func GetToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	ts, err := token.SignedString(signingKey)
	return ts, err
}

func VerifyToken(ts string) (jwt.Claims, error) {
	token, err := jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
