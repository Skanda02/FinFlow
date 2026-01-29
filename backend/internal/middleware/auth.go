package middleware

import (
	"context"
	"net/http"
	"strings"

	"finflow/internal/http_helpers"
	"finflow/internal/utils"
)

type contextKey string

const UserIDKey contextKey = "user_id"

// AuthMiddleware validates JWT token and adds user ID to request context
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http_helpers.WriteJSONError(w, http.StatusUnauthorized, "authorization header required")
			return
		}

		// Extract token from "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http_helpers.WriteJSONError(w, http.StatusUnauthorized, "invalid authorization header format")
			return
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			http_helpers.WriteJSONError(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		// Add user ID to context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// GetUserIDFromContext retrieves the user ID from request context
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(UserIDKey).(int)
	return userID, ok
}
