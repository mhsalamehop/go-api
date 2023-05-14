package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey = []byte(os.Getenv("JWT_SECRET"))
)

func GetToken(email,role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role": role,
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

// `INSERT INTO table_name clms` values() ON CONFLICT(pk) DO UPDATE SET id=$1 -- movies backfill