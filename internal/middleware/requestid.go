package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const RequestIDKey contextKey = "requestID"

// RequestID middleware adds a unique request ID to each request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if request ID already exists in headers
		requestID := r.Header.Get("X-Request-ID")

		// If not, generate a new one
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Add to response headers
		w.Header().Set("X-Request-ID", requestID)

		// Add to context for use in handlers
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)

		// Call the next handler with updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID retrieves the request ID from the context
func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}
