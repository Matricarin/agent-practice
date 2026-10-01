package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type sumRequest struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type sumResponse struct {
	Sum float64 `json:"sum"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, errorResponse{Error: msg})
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req sumRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.A == nil || req.B == nil {
		writeError(w, http.StatusBadRequest, "fields a and b are required numbers")
		return
	}

	writeJSON(w, http.StatusOK, sumResponse{Sum: *req.A + *req.B})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sum", sumHandler)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
