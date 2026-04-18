package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()

	metricsHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status esperado 200, obtido %d", resp.StatusCode)
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("Content-Type esperado application/json, obtido %s", ct)
	}
}