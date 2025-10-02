package handler

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey contextKey = "user"

// contextWithUser saves JWT claims in request context
func contextWithUser(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, userContextKey, claims)
}

// GetUserFromContext retrieves JWT claims from request context
func GetUserFromContext(r *http.Request) (jwt.MapClaims, bool) {
	claims, ok := r.Context().Value(userContextKey).(jwt.MapClaims)
	return claims, ok
}
