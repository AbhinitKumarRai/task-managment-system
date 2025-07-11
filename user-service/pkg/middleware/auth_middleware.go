// middleware/auth.go
package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid Authorization header", http.StatusUnauthorized)
			return
		}
		token := parts[1]
		if !strings.HasPrefix(token, "user-") {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}
		userIDStr := strings.TrimPrefix(token, "user-")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "invalid user id", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
