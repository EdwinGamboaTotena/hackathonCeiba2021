package configuraciones

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func IniciarAplicacion() {
	router := mux.NewRouter()
	definirRutas(router)

	puerto := ":5000"

	servidor := &http.Server{
		Handler:      router,
		Addr:         puerto,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Escuchando en %s\n", puerto)
	log.Fatal(servidor.ListenAndServe())
}
