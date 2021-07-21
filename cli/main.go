package main

import (
	"flag"
	"golim/internal/middleware"
	"golim/internal/proxy"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/time/rate"
)

var (
	webServerAddr            = "http://127.0.0.1:8000"
	serverAddress            = "localhost:8080"
	rateLimiterBurst         = 10
	rateLimiterWindowSeconds = 1
)

func getProxyHandler(target *url.URL) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.HTTPRequest(target, w, r)
	}
}

func lookupEnvOrDefault(key, def interface{}) interface{} {
	return def
}

func main() {
	// set up logging
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// set up flags
	flag.StringVar(&serverAddress, "server-addr", lookupEnvOrDefault("GOLIM_SERVER_ADDR", serverAddress).(string), "address of this server")
	flag.StringVar(&webServerAddr, "web-server-addr", lookupEnvOrDefault("GOLIM_WEB_SERVER_ADDR", webServerAddr).(string), "address of the web server")
	flag.IntVar(&rateLimiterBurst, "rate-limiter-burst", lookupEnvOrDefault("GOLIM_RATE_LIMITER_BURST", rateLimiterBurst).(int), "number of requests that can be made in a given time window")
	flag.IntVar(&rateLimiterWindowSeconds, "rate-limiter-window-seconds", lookupEnvOrDefault("GOLIM_RATE_LIMITER_WINDOW_SECONDS", rateLimiterWindowSeconds).(int), "size of the rate limiter window")
	flag.Parse()

	webServerAddrURL, err := url.Parse(webServerAddr)
	if err != nil {
		logger.Fatalf("Error parsing the web server address as URL: %v", err)
	}

	// set up the http routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", getProxyHandler(webServerAddrURL))

	// set up the middlewares
	middlewareMux := middleware.NewRateLimiterMiddleware(logger, rate.Every(time.Duration(rateLimiterWindowSeconds)*time.Second), rateLimiterBurst, mux)
	middlewareMux = middleware.NewHTTPLoggerMiddleware(logger, middlewareMux)
	middlewareMux = middleware.NewRequestIDMiddleware(middlewareMux)

	// start the server
	logger.Printf("Starting the server on %s\n", serverAddress)

	err = http.ListenAndServe(serverAddress, middlewareMux)
	if err != nil {
		logger.Fatalf("Error starting the server: %v\n", err)
	}
}
