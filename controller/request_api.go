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
	_, err := strconv.Atoi(parameter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "BadRequest",
			"error":  "The QueryParam must be an Integer",
		})
		return
	}

	service := services.GetInstance()
	response, err := service.RequestApi(parameter)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ServiceUnavailable",
			"error":  "The external api is not available, try again later ",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
