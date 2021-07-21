package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey string

const ctxKeyRequestID ctxKey = "requestId"

type requestIDHandler struct {
	next http.Handler
}

func setRequestID(ctx context.Context) context.Context {
	id := uuid.New()
	return context.WithValue(ctx, ctxKeyRequestID, id.String())
}

// GetRequestID fetches and returns the request ID from the context.
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(ctxKeyRequestID).(string); ok {
		return id
	}
	return ""
}

// ServeHTTP binds requestIDHandler to the Handler interface in net/http.
// It adds a request ID to the request context and response header.
func (h requestIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// set the request ID in the request context
	ctx := setRequestID(r.Context())
	r = r.WithContext(ctx)

	// set the request ID in the response header
	w.Header().Set("X-Request-ID", GetRequestID(ctx))

	// call the next handler
	h.next.ServeHTTP(w, r)
}

// NewRequestIDHandler gets the request ID middleware
func NewRequestIDHandler(next http.Handler) http.Handler {
	return requestIDHandler{next}
}
