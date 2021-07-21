package main

import (
	"log"
	"net/http"

	"golim/internal/middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.Handle("/", middleware.RequestIDHandler(http.HandlerFunc(handler)))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error starting the server: %v\n", err)
	}
}
