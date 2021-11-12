package config

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	configureMapping(router)

	port := ":5000"

	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Escuchando en %s\n", port)
	log.Fatal(server.ListenAndServe())
}
