package config

import (
	"github.com/EdwinGamboaTotena/hackathonCeiba2021/controller"
	"github.com/gorilla/mux"
)

func configureMapping(router *mux.Router) {
	router.HandleFunc("/healthz", controller.Healthz).Methods("GET")
	router.HandleFunc("/", controller.RequestApi).Methods("GET")
}
