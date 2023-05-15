package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mhsalamehop/go-api/utils"
)

type Authorization struct{}

func (amw *Authorization) IsAuthorized(userRole string,next http.Handler) http.Handler {
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
			return
		}
		email, ok := claims.(jwt.MapClaims)["email"].(string)
		if !ok {
			http.Error(w, "invalid role claim in token", http.StatusUnauthorized)
			return
		}
		if !validEmail(email){
			http.Error(w, "invalid email format", http.StatusUnauthorized)
			return
		}
		role, ok := claims.(jwt.MapClaims)["role"].(string)
		if !ok {
			http.Error(w, "invalid role claim in token", http.StatusUnauthorized)
			return
		}
		if role != userRole {
			http.Error(w,"Access Denied",http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validEmail(email string) bool {
	return strings.HasSuffix(email, "@gmail.com")
}