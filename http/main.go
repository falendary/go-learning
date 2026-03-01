package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Echo string `json:"echo"`
	Time string `json:"time"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		w.Header().Set("Allow", http.MethodGet)
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	resp := HealthResponse{
		Status: "ok",
		Time:   time.Now().Format(time.RFC3339),
	}

	writeJSON(w, http.StatusOK, resp)

}

func readJSON(r *http.Request, dst any) error {

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // helps catch typos in JSON keys
	return dec.Decode(dst)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return

	}

	var req EchoRequest

	if err := readJSON(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}

	if req.Message == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "message is required"})
	}

	writeJSON(w, http.StatusOK, EchoResponse{Echo: req.Message, Time: time.Now().Format(time.RFC3339)})
}

func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %s", r.Method, r.URL.Path, time.Since(start))
	})

}

func main() {

	mux := http.NewServeMux()

	// Routes

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/echo", echoHandler)

	handler := loggingMiddleware(mux)

	src := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Listening on http://localhost:8080")
	log.Fatal(src.ListenAndServe())
}
