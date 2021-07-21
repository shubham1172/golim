package middleware

import (
	"log"
	"net/http"
	"time"
)

type httpLoggerHandler struct {
	logger *log.Logger
	next   http.Handler
}

// ServeHTTP binds httpLoggerHandler to the Handler interface in net/http.
// It logs details about the incoming request and outgoing response.
func (h httpLoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := GetRequestID(r.Context())

	h.logger.Printf("[%s] Incoming %s %s %s", id, r.Method, r.URL, r.RemoteAddr)

	startTime := time.Now()
	h.next.ServeHTTP(w, r)
	duration := time.Since(startTime)

	h.logger.Printf("[%s] Finished %s %s %s %dms", id, r.Method, r.URL, r.RemoteAddr, duration.Milliseconds())
}

// NewHTTPLoggerMiddleware gets the logger middleware.
func NewHTTPLoggerMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return httpLoggerHandler{logger, next}
}
