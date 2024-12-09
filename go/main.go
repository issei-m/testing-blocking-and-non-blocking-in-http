package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

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

func handler(w http.ResponseWriter, r *http.Request) {
	sleep := r.URL.Query().Get("sleep")
	targetURL := fmt.Sprintf("http://rust:3000/?sleep=%s", url.QueryEscape(sleep))

	w.Header().Set("Content-Type", "text/plain")

	resp, err := http.Get(targetURL)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	loggedMux := logger(mux)

	fmt.Println("Starting server at :8000")
	if err := http.ListenAndServe(":8000", loggedMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
