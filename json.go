package main

import (
	"log"

	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{Error: message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
