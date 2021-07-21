package main

import (
	"log"
	"net/http"
	"os"

	"golim/internal/middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	logger := log.New(os.Stdout, "", log.LstdFlags)
	m := middleware.NewRequestIDHandler(middleware.NewLoggerMiddleware(logger, mux))

	addr := "localhost:8080"
	logger.Printf("Starting the server on %s\n", addr)

	err := http.ListenAndServe(addr, m)
	if err != nil {
		logger.Fatalf("Error starting the server: %v\n", err)
	}
}
