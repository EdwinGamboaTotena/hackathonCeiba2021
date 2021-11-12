package configuraciones

import (
	"github.com/EdwinGamboaTotena/hackathonCeiba2021/controladores"
	"github.com/gorilla/mux"
)

func definirRutas(router *mux.Router) {
	router.HandleFunc("/healthz", controladores.StatusControlador).Methods("GET")
	router.HandleFunc("/", controladores.ConsultaApiControlador).Methods("GET")
}
