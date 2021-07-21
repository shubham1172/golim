package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// HTTPRequest invokes target URL and returns its response.
func HTTPRequest(target *url.URL, w http.ResponseWriter, r *http.Request) {
	r.URL.Host = target.Host
	r.URL.Scheme = target.Scheme
	r.Host = target.Host
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Header.Set("X-Origin-Host", r.Host)

	// proxy it to the target
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}
