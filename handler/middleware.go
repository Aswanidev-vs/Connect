package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

// JWTMiddleware protects routes by validating the JWT token in cookies
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			// No token → redirect to login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return JwtKey, nil
		})

		if err != nil || !token.Valid {
			// Invalid token → redirect to login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// (Optional) save claims in context for downstream handlers
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx := contextWithUser(r.Context(), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// If claims are not valid → force login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
