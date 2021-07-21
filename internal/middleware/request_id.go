package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey string

const ctxKeyRequestID ctxKey = "requestId"

func setRequestID(ctx context.Context) context.Context {
	id := uuid.New()
	return context.WithValue(ctx, ctxKeyRequestID, id.String())
}

// GetRequestID fetches and returns the request ID from the context
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(ctxKeyRequestID).(string); ok {
		return id
	}
	return ""
}

// RequestIDHandler adds a request ID to the request context and response header
func RequestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

		ctx := setRequestID(r.Context())
		r = r.WithContext(ctx)
		w.Header().Set("X-Request-ID", GetRequestID(ctx))
	})
}
