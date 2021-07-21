package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggerHandler struct {
	logger *log.Logger
	next   http.Handler
}

// ServeHTTP binds loggerHandler to the Handler interface in net/http.
// It logs details about the incoming request and outgoing response.
func (h loggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := GetRequestID(r.Context())

	h.logger.Printf("[%s] Incoming %s %s %s", id, r.Method, r.URL, r.RemoteAddr)

	startTime := time.Now()
	h.next.ServeHTTP(w, r)
	duration := time.Since(startTime)

	h.logger.Printf("[%s] Finished %s %s %s %dms", id, r.Method, r.URL, r.RemoteAddr, duration.Milliseconds())
}

// NewLoggerMiddleware gets the logger middleware.
func NewLoggerMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return loggerHandler{logger, next}
}
