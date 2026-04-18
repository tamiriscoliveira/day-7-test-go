package main

import (
	"encoding/json"
	"net/http"
)

type Metrics struct {
	Memory float64 `json:"memory"`
	CPU    float64 `json:"cpu"`
	Disk   float64 `json:"disk"`
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := Metrics{
		Memory: 42.0,
		CPU:    27.3,
		Disk:   63.9,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(":8080", nil)
}
