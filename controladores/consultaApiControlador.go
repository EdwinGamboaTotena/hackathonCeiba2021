package controladores

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EdwinGamboaTotena/hackathonCeiba2021/servicios"
)

func ConsultaApiControlador(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parametro := r.FormValue("number")
	numero, err := strconv.Atoi(parametro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respuesta, err := servicios.ConsultarApi(numero)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
