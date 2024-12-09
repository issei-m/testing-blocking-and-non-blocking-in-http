package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crw := &customResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(crw, r)

		log.Printf("- %s \"%s %s\" %d",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			crw.statusCode,
		)
	})
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

func handler(w http.ResponseWriter, r *http.Request) {
	sleepParam := r.URL.Query().Get("sleep")
	sleepTime, err := strconv.Atoi(sleepParam)
	if err != nil || sleepTime < 0 {
		sleepTime = 10 // Default to 10 seconds if invalid
	}

	time.Sleep(time.Duration(sleepTime) * time.Second)

	fmt.Fprintf(w, "%d seconds have passed", sleepTime)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	loggedMux := logger(mux)

	fmt.Println("Starting server at :8888")
	if err := http.ListenAndServe(":8888", loggedMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
