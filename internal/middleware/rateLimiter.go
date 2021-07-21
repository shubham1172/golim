package middleware

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

type rateLimiterHandler struct {
	next    http.Handler
	limiter *rate.Limiter
	logger  *log.Logger
}

// ServeHTTP binds rateLimiterHandler to the Handler interface in net/http.
// If number of requests exceed the rate limit, it returns a HTTP 429 response.
func (h rateLimiterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := GetRequestID(r.Context())

	if !h.limiter.Allow() {
		h.logger.Printf("[%s] Request throttled by rate-limiter.", id)
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}
	h.next.ServeHTTP(w, r)
}

// NewRateLimiterMiddleware gets the rate limiter middleware.
// It allows b requests at a rate of r.
// For example, to allow 10 requests per second, use
// 	`NewRateLimiterMiddleware(rate.Every(time.Second), 10, next)`
func NewRateLimiterMiddleware(logger *log.Logger, r rate.Limit, b int, next http.Handler) http.Handler {
	limiter := rate.NewLimiter(r, b)
	return rateLimiterHandler{next, limiter, logger}
}
