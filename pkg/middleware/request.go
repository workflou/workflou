package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := uuid.New().String()
			ctx := context.WithValue(r.Context(), "requestID", id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
