package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Service: "Nebula Controller",
		Status:  "running",
		Version: "0.1.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", rootHandler)

	log.Println("Nebula Controller running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}