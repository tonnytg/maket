package webserver

import (
	"github.com/google/uuid"
	"log"
	"net/http"
)

func middleWareChecker(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.Header.Set("X-Request-ID", uuid.New().String())

		next(w, r)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func Start() {

	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/api/v1/health", middleWareChecker(healthCheck))

	// Add new endpoint
	muxServer.HandleFunc("/api/v1/targets", middleWareChecker(handleTarget))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", muxServer); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
