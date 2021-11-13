package controller

import (
	"encoding/json"
	"net/http"
)

func Healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "version": "1.1.1"})
}
