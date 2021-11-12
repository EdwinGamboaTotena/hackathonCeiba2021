package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EdwinGamboaTotena/hackathonCeiba2021/services"
)

func RequestApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parameter := r.FormValue("number")
	number, err := strconv.Atoi(parameter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := services.RequestApi(number)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
