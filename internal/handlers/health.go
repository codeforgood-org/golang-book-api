package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthCheck handles health check requests
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}
