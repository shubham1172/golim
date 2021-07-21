package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"golim/internal/middleware"

	"golang.org/x/time/rate"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	middlewareMux := middleware.NewRateLimiterMiddleware(logger, rate.Every(time.Second), 1, mux)
	middlewareMux = middleware.NewHTTPLoggerMiddleware(logger, middlewareMux)
	middlewareMux = middleware.NewRequestIDMiddleware(middlewareMux)

	addr := "localhost:8080"
	logger.Printf("Starting the server on %s\n", addr)

	err := http.ListenAndServe(addr, middlewareMux)
	if err != nil {
		logger.Fatalf("Error starting the server: %v\n", err)
	}
}
