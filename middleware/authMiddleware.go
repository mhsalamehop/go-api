package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mhsalamehop/go-api/utils"
)

type Authorization struct{}

func (amw *Authorization) IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := r.Header.Get("Authorization")
		if len(ts) == 0 {
			http.Error(w, "must provide auth token", http.StatusUnauthorized)
			return
		}
		ts = strings.Split(ts, " ")[1]
		claims, err := utils.VerifyToken(ts)
		if err != nil {
			http.Error(w, "error verifying JWT token "+err.Error(), http.StatusUnauthorized)
		}
		email := claims.(jwt.MapClaims)["email"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)
		r.Header.Set("email", email)
		r.Header.Set("role", role)
		next.ServeHTTP(w, r)
	})
}
